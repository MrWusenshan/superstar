package conf

const DriverName = "mysql"

type DbConfig struct {
	Host   string
	Port   int
	User   string
	Psd    string
	DbName string
}

var MasterDbConfig DbConfig = DbConfig{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Psd:    "root",
	DbName: "superstar",
}

var SlaveDbConfig DbConfig = DbConfig{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Psd:    "root",
	DbName: "superstar",
}
