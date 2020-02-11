package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go-blog/models/admin"
	"go-blog/utils"
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
	count,_ := nqs.Count()
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
	c.Data["DateCount"] = count

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
	reviewCount ,_ := nqrw.Count()
	c.Data["ReviewCount"] = reviewCount
    c.Data["Review"] = reviewData
}

func (c *BaseController) Menu()  {

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	menu , _  := admin.GetAllMenu(query,fields,sortby,order,offset,limit)
	data := utils.MenuData(menu,0,0)
	/*c.Data["json"] = data
	c.ServeJSON()
	c.StopRun()*/
	c.Data["Menu"] = data

	link , _  := admin.GetAllLink(query,fields,sortby,order,offset,limit)
	c.Data["Link"] = link

}
func (c *BaseController) Prepare(){
	c.Data["bgClass"] = "bgColor"
	c.Data["T"] = time.Now()

	o := orm.NewOrm()
	var setting []*admin.Setting
	o.QueryTable(new(admin.Setting)).All(&setting)

	for _,v := range setting{
		c.Data[v.Name] = v.Value
	}

	pv,_ := o.QueryTable(new(admin.Log)).Count()
	uv,_ := o.QueryTable(new(admin.Log)).GroupBy("ip").Count()

	c.Data["PV"] = pv
	c.Data["UV"] = uv
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


