package dao

import (
	"log"
	"superstarProject/models"
	"github.com/go-xorm/xorm"
)

type SuperstarDao struct {
	engine *xorm.Engine
}

func NewSuperstarDao(engine *xorm.Engine) *SuperstarDao {
	return &SuperstarDao{engine: engine}
}

//根据id查单条数据
func (d *SuperstarDao) Get(id int) *models.StarInfo {
	data := &models.StarInfo{}

	ok, err := d.engine.Get(id)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0 //没有查找到数据或者查询出错，将id=0的数据返回
		return data
		//return nil    或者直接返回nil
	}
}

//查询全部数据
func (d *SuperstarDao) GetAll() []models.StarInfo {
	datalist := []models.StarInfo{} //或者	datalist:=make([]models.StarInfo.0)

	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		//查询出错处理
		log.Println(err)
		return datalist //或者直接返回nil 		return nil
	}

	return datalist
}

//根据 id 删除数据
func (d *SuperstarDao) Delete(id int) error {
	data := &models.StarInfo{Id: id, SysStatus: 1}

	_, err := d.engine.Id(data.Id).Update(data)
	return err

	//或者这种写法
	//_,err:=d.engine.Delete(id)
	//if err==nil {
	//	return true
	//}
	//return false
}

//coiumns 是否强制更新
func (d *SuperstarDao) Updata(data *models.StarInfo, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

//添加数据
func (d *SuperstarDao) Create(data *models.StarInfo) error {
	_, err := d.engine.Insert(data)
	return err
}

//通过 country 查询数据
func (d *SuperstarDao) Serch(country string) []models.StarInfo {
	datalist:=[]models.StarInfo{}		//或者	datalist:=make([]models.StarInfo.0)

	err:=d.engine.Where("country=?",country).Desc("id").Find(&datalist)
	if err!=nil{
		log.Println(err)	//或者直接返回空   return nil
		return datalist
	}
	return datalist
}