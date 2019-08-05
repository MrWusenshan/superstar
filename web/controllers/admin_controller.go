package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"superstarProject/models"
	"superstarProject/services"
	"time"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

//admin 首页
func (c *AdminController) GetAll() mvc.Result {
	datalist := c.Service.GetAll()
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "球星库后台管理",
			"Datalist": datalist,
		},
		Layout: "admin/layout.html",
	}
}

//查单个数据
func (c *AdminController) GetEdit() mvc.Result {
	var data *models.StarInfo

	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		data = c.Service.Get(id)
	}

	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title": "球星库后台管理",
			"Data":  data,
		},
	}
}

//修改
func (c *AdminController) PostSave() mvc.Result {
	info := models.StarInfo{}
	err := c.Ctx.ReadForm(&info)
	if err != nil {
		log.Fatal(err)
	}

	if info.Id > 0 {
		info.SysUpdated = int(time.Now().Unix())
		c.Service.Update(&info, []string{"name_zh", "name_en", "avatar",
			"birthday", "height", "weight", "club", "jersy", "coutry",
			"birthaddress", "feature", "moreinfo", "sys_updated"})
	} else {
		info.SysStatus = int(time.Now().Unix())
		c.Service.Create(&info)
	}

	return mvc.Response{
		Path: "/admim/",
	}
}

//删除
func (c *AdminController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}
	return mvc.Response{
		Path:"/admin/",
	}
}
