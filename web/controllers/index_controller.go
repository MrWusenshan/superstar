// file: controllers/index_controller.go

package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"superstar/services"
)

type IndexController struct {
	Ctx      iris.Context
	Services services.SuperstarService
}

//首页    method:GET    获取所有球星数据
func (c *IndexController) GetIndex() mvc.Result {
	datalist:=c.Services.GetAll()
	return mvc.View{
		Name:"index.html",
		Data:iris.Map{
			"Title":"球星库",
			"Datalist":datalist,
		},
	}
}

//
func (c *IndexController) GetBy(id int) mvc.Result {
	//id小于1 错误 返回首页
	if id<1{
		return mvc.Response{
			Path:"/",
		}
	}

	data:=c.Services.Get(id)
	return mvc.View{
		Name:"info.html",
		Data:iris.Map{
			"Title":"球星库",
			"Data":data,
		},
	}
}

func (c *IndexController) GetSerch() mvc.Result{
	country:=c.Ctx.URLParam("country")
	datalist:=c.Services.Serch(country)
	return mvc.View{
		Name:"index.html",
		Data:iris.Map{
			"Title":"球星库",
			"Datalist":datalist,
		},
	}
}

