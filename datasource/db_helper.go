package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"superstarProject/conf"
	"sync"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

//master 	链接mysql数据库
func InstanceMaster() *xorm.Engine {
	//第一次判断正常执行
	if masterEngine != nil {
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	//第二次判断,防止多个请求多次初始化
	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConfig

	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)%s?charset=utf8",
		c.User, c.Psd, c.Host, c.Port, c.DbName)

	engine, err := xorm.NewEngine(conf.DriverName, driverSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster:", err)
		return nil
	}
	return engine
}

//slave 链接mysql数据库
func InstanceSlave() *xorm.Engine {
	//第一次判断正常执行
	if slaveEngine != nil {
		return slaveEngine
	}

	lock.Lock()
	defer lock.Unlock()

	//第二次判断,防止多个请求多次初始化
	if slaveEngine != nil {
		return slaveEngine
	}

	c := conf.SlaveDbConfig

	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)%s?charset=utf8",
		c.User, c.Psd, c.Host, c.Port, c.DbName)

	engine, err := xorm.NewEngine(conf.DriverName, driverSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstancSlavce:", err)
		return nil
	}
	return engine
}
