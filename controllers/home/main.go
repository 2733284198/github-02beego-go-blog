package home

import "github.com/astaxie/beego"

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Log("index")
	c.TplName = "home/" + beego.AppConfig.String("view") + "/index.html"
}


