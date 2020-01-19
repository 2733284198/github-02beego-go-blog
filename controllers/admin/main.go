package admin

import (
	"github.com/astaxie/beego/orm"
	"go-blog/utils"
	"time"
)


type Log struct {
	Num 	int
	Date 	string
	Ip 		string
}

type MainController struct {
	BaseController
}

func (c *MainController) Index() {
	c.TplName = "admin/index.html"
}

func (c *MainController) Welcome() {


	o := orm.NewOrm()

	var log []orm.Params

	// time.Now().AddDate(0,0,-7) 七天前
	o.Raw("SELECT `id`,`ip`,`create` FROM `log` where `create` >= ? ORDER BY `create` ASC",time.Now().AddDate(0,0,-7)).Values(&log)

	var pv = make(map[string]Log)
	var uv = make(map[string]Log)
	var dateSlice []string
	var pvSlice []int
	var uvSlice []int

	for _,v := range log {
		//utils.StringToTime(v["create"])
		// 获取日期
		var key = utils.StringToTime(v["create"]).Format("2006-01-02")

		// 统计pv
		pvTemp := pv[key]
		pv[key] = Log{
			Num:pvTemp.Num + 1,
		}
		// 统计uv
		uvTemp := uv[key]
		if uvTemp.Ip != v["ip"].(string) {
			uv[key] = Log{
				Num:uvTemp.Num + 1,
			}
		}

	}


	for k,v := range pv {
		dateSlice = append(dateSlice,k)
		pvSlice = append(pvSlice,v.Num)
	}

	for _,v := range uv {
		uvSlice = append(uvSlice,v.Num)
	}

	c.Data["Date"] = dateSlice
	c.Data["Pv"] = pvSlice
	c.Data["Uv"] = uvSlice

	c.TplName = "admin/welcome.html"
}
