package routes

import (
	"github.com/kataras/iris/mvc"
	"superstarProject/services"
	"superstarProject/bootstrap"
	"superstarProject/web/middleware"
	"superstarProject/web/controllers"
)

func Configure(b *bootstrap.Bootstrapper) {
	superstarService := services.NewSupertarService()

	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controllers.AdminController))
}
