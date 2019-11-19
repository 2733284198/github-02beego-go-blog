package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"leechan.inline/models/admin"
	"time"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Layout() {

	o := orm.NewOrm()

	// 月份排序
	articleTime := new(admin.Article)
	var articlesTime []*admin.Article
	nqs := o.QueryTable(articleTime)
	nqs = nqs.Filter("status", 1)
	nqs.OrderBy("-Created").RelatedSel().All(&articlesTime,"Created")
	var datetime = make(map[string]int64)
	var dateTimeKey []string
	for _,v := range articlesTime  {
		//str = append(str ,v.Created.Format("2006-01"))
		//c.Ctx.WriteString(v.Created.Format("2006-01"))
		k := v.Created.Format("2006-01")
		if datetime[k] == 0{
			dateTimeKey = append(dateTimeKey,k)
		}
		datetime[k] = datetime[k] + 1
	}
	c.Data["DateTime"] = datetime
	c.Data["DateTimeKey"] = dateTimeKey

	// 阅读排序
	articleReadSort := new(admin.Article)
	var articlesReadSort []*admin.Article
	nqrs := o.QueryTable(articleReadSort)
	nqrs = nqrs.Filter("status", 1)
	nqrs = nqrs.OrderBy("-Pv")
	nqrs.Limit(5).All(&articlesReadSort,"Id","Title","Pv")
	c.Data["ArticlesReadSort"] = articlesReadSort
	
	// 最新评论
        review := new(admin.Review)
        var reviewData []*admin.Review
        nqrw := o.QueryTable(review)
        nqrw = nqrw.Filter("status", 1)
        nqrw = nqrw.OrderBy("-Id")
        nqrw.Limit(5).All(&reviewData,"Review","ArticleId")
        c.Data["Review"] = reviewData
}

func (c *BaseController) Menu()  {

	o := orm.NewOrm()

	category := new(admin.Category)
	var categorys []*admin.Category
	cqs := o.QueryTable(category)
	cqs = cqs.Filter("status", 1)
	cqs = cqs.Filter("pid", 0)
	cqs.OrderBy("-sort").All(&categorys)
	c.Data["Menu"] = categorys

}
func (c *BaseController) Prepare(){
	c.Data["bgClass"] = "bgColor"
	c.Data["T"] = time.Now()
}

/*********************** 日志 *********************************/

type LogData struct {
	Code		int   		`json:"code"`
	Ip		Ip       	`json:"data"`
}

type Ip struct {
	Area              string   `json:"area"`
	City              string   `json:"city"`
	Region			  string   `json:"region"`
}

func (c *BaseController) Log(page string)  {

	ip := c.Ctx.Input.IP()
	/*
	resp, err := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip="+ip)

	var city string
	if err == nil{

		readAll, err := ioutil.ReadAll(resp.Body)
		if err == nil{
			var data LogData
			json.Unmarshal(readAll, &data)
			city = data.Ip.City
		}

	}

	defer resp.Body.Close()
	*/
	userAgent := c.Ctx.Input.UserAgent()

	url := c.Ctx.Input.URI()
	o := orm.NewOrm()
	var log = admin.Log{
		Ip:    			ip,
		//City:     		city,
		UserAgent:    	userAgent,
		Page:			page,
		Uri:			url,
	}
	o.Insert(&log)

}


