package admin

import (
	"go-blog/utils/article"
	"go-blog/utils"
	"github.com/astaxie/beego/orm"
	"go-blog/models/admin"
	"time"
)

type ArticleResourcesController struct {
	BaseController
}

func (c *ArticleResourcesController) GetArticleList()  {
	/*gocn := Gocn{}
	title,html := gocn.Get("")
	md := utils.Html2md(html)

	o := orm.NewOrm()
	article := admin.Article{
		Title:    title,
		Tag:      tag,
		Desc:     md,
		Html:	  html,
		Remark:   remark,
		Status:   1,
		User:     &admin.User{1, "", "", "", time.Now(), 0},
		Category: &admin.Category{13, "", 0, 0, 0},
	}

	response := make(map[string]interface{})

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			//log.Println(err.Key, err.Message)
			response["msg"] = "新增失败！"
			response["code"] = 500
			response["err"] = err.Key + " " + err.Message
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}
	}



	if id, err := o.Insert(&article); err == nil {
		response["msg"] = "新增成功！"
		response["code"] = 200
		response["id"] = id
	} else {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()*/

	c.TplName = "admin/article-get.html"
}

func(c *ArticleResourcesController)GetArticle(){

	response := make(map[string]interface{})
	url := c.GetString("url")
	if url == ""{
		response["msg"] = "抓取失败！"
		response["code"] = 500
		response["err"] = "url must no be null"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	gocn := article.Gocn{}
	title,html := gocn.Get(url)
	md := utils.Html2md(html)

	o := orm.NewOrm()
	article := admin.Article{
		Title:    title,
		Tag:      "",
		Desc:     md,
		Html:	  html,
		Remark:   "",
		Status:   1,
		User:     &admin.User{1, "", "", "", time.Now(), 0},
		Category: &admin.Category{13, "", 0, 0, 0},
	}
	if id, err := o.Insert(&article); err == nil {
		response["msg"] = "新增成功！"
		response["code"] = 200
		response["id"] = id
	} else {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()

	
}