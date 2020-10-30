package admin

import (
	"go-blog/utils/article"
	"go-blog/utils"
	"github.com/astaxie/beego/orm"
	"go-blog/models/admin"
	"time"
	"fmt"
	"net/url"
	"strings"
)

type ArticleResourcesController struct {
	BaseController
}

func(c *ArticleResourcesController)GetArticle(){

	t1 := time.Now()
	response := make(map[string]interface{})
	urlPath := c.GetString("url")
	if urlPath == ""{
		response["msg"] = "抓取失败！"
		response["code"] = 500
		response["err"] = "url must no be null"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}


	u, err := url.Parse(urlPath)
	if err != nil {
        response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
    }
	host := u.Host
	ho := strings.Split(host, ":")

	var title,html,md string
	switch ho[0]{
		case "gocn.vip" :
			gocn := article.Gocn{}
			title,html = gocn.Get(urlPath)
			md = utils.Html2md(html)
		default:
			response["msg"] = "新增失败！"
			response["code"] = 500
			response["err"] = "未知源!"
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
	}

	
	o := orm.NewOrm()
	article := admin.Article{
		Title:    title,
		Tag:      "",
		Desc:     md,
		Html:	  html,
		Remark:   "",
		Status:   1,
		User:     &admin.User{1, "", "", "", time.Now(), 0},
		Category: &admin.Category{30, "", 0, 0, 0},
	}

	if id, err := o.Insert(&article); err == nil {
		response["msg"] = "新增成功！"
		response["code"] = 200
		response["id"] = id
		response["title"] = title
		response["elapsed"] = fmt.Sprintf("%s",time.Since(t1))
	} else {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func(c *ArticleResourcesController)GetCron(){
	c.TplName = "admin/article-auto-add.html"
}