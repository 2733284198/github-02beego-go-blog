package home

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.TplName = "home/index.html"
}


