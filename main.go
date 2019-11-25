package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	_ "leechan.inline/routers"
	"leechan.inline/utils"
)

var TplFuncMap = make(template.FuncMap)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "xxxx:xxxx@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Asia%2FShanghai")

	beego.AddFuncMap("IndexForOne", utils.IndexForOne)
	beego.AddFuncMap("IndexAddOne",utils.IndexAddOne)
	beego.AddFuncMap("IndexDecrOne",utils.IndexDecrOne)
	beego.AddFuncMap("StringReplace",utils.StringReplace)

}

func main() {
	beego.Run()
}
