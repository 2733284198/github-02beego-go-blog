package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"html/template"
	_ "go-blog/routers"
	"go-blog/utils"
)

var TplFuncMap = make(template.FuncMap)
func init() {
	conf, err := config.NewConfig("ini", "conf/app.conf")

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	dbUser := conf.String("db::dbUser")
	dbPass := conf.String("db::dbPass")
	dbHost := conf.String("db::dbHost")
	dbPort := conf.String("db::dbPort")
	dbName := conf.String("db::dbName")
	dbStr := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName+ "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", dbStr)

	beego.AddFuncMap("IndexForOne", utils.IndexForOne)
	beego.AddFuncMap("IndexAddOne",utils.IndexAddOne)
	beego.AddFuncMap("IndexDecrOne",utils.IndexDecrOne)
	beego.AddFuncMap("StringReplace",utils.StringReplace)
	beego.AddFuncMap("TimeStampToTime",utils.TimeStampToTime)

}

func main() {
	beego.Run()
}
