package home


import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go-blog/models/admin"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Log("index")

	// 推荐
	o := orm.NewOrm()
	var list []*admin.Article
	o.QueryTable(new(admin.Article)).Filter("status", 1).Filter("recommend", 1).Filter("User__Name__isnull", false).Filter("Category__Name__isnull", false).OrderBy("-id").RelatedSel().All(&list, "id", "title")
	c.Data["Recommend"] = list

	c.TplName = "home/" + beego.AppConfig.String("view") + "/index.html"
}


