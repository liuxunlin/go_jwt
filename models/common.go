package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost := beego.AppConfig.String("mysql::dbhost")
	dbport := beego.AppConfig.String("mysql::dbport")
	dbuser := beego.AppConfig.String("mysql::dbuser")
	dbpassword := beego.AppConfig.String("mysql::dbpassword")
	dbname := beego.AppConfig.String("mysql::dbname")
	if dbport == "" {
		dbport = "13306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"

	err := orm.RegisterDataBase("default", "mysql", dsn, 30)
	if err != nil {
		fmt.Println("mysql connection error: ", err)
	}
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("mysql::dbprefix") + str
}
