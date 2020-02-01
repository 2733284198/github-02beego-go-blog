package home

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Log("index")
	c.TplName = "home/index.html"
}


