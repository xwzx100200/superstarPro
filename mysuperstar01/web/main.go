package main

import (
	"superstartPro/superstar/bootstrap"
	"superstartPro/superstar/web/middleware/identity"
	"superstartPro/superstar/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "一凡Sir")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
