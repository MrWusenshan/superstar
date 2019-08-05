package main

import (
	"superstarProject/routes"
	"superstarProject/bootstrap"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "Quincy")
	app.Bootstrap()
	app.Configure( routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
