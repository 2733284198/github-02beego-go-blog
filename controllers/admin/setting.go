package admin

import (
	"github.com/astaxie/beego/orm"
	"go-blog/models/admin"
)

type SettingController struct {
	BaseController
}

func (c *SettingController) Add() {

	o := orm.NewOrm()
	var setting []*admin.Setting
	o.QueryTable(new(admin.Setting)).All(&setting)

	for _,v := range setting{
		c.Data[v.Name] = v.Value
	}

	c.TplName = "admin/setting.html"
}

func (c *SettingController) Save() {

	response := make(map[string]interface{})

	title := c.GetString("title")
	name := c.GetString("name")
	tag := c.GetString("tag")
	remark := c.GetString("remark")
	image := c.GetString("image")


	o := orm.NewOrm()
	err := o.Begin()


	_,err = o.Delete(&admin.Setting{Name: "name"})
	_,err = o.Delete(&admin.Setting{Name: "title"})
	_,err = o.Delete(&admin.Setting{Name: "tag"})
	_,err = o.Delete(&admin.Setting{Name: "remark"})
	_,err = o.Delete(&admin.Setting{Name: "image"})

	settings := []admin.Setting{
		{Name: "title", Value: title},
		{Name: "name", Value: name},
		{Name: "tag",Value:tag},
		{Name: "remark",Value:remark},
		{Name: "image",Value:image},
	}

	num, err := o.InsertMulti(5, settings)

	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	if err != nil {
		response["msg"] = "操作失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}else{
		response["msg"] = "操作成功！"
		response["code"] = 200
		response["id"] = num
	}


	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func (c *SettingController) Notice() {

	o := orm.NewOrm()
	var setting admin.Setting
	o.QueryTable(new(admin.Setting)).Filter("name", "notice").One(&setting)
	c.Data["Notice"] = setting.Value
	c.TplName = "admin/notice.html"
}

func (c *SettingController) NoticeSave() {

	response := make(map[string]interface{})

	notice := c.GetString("notice")

	o := orm.NewOrm()
	err := o.Begin()
	_,err = o.Delete(&admin.Setting{Name: "notice"})

	num, err := o.Insert(&admin.Setting{Name:"notice",Value:notice})

	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}

	if err != nil {
		response["msg"] = "操作失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}else{
		response["msg"] = "操作成功！"
		response["code"] = 200
		response["id"] = num
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}
