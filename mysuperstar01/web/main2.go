package main

import (
	"superstartPro/superstar/bootstrap"
	"superstartPro/superstar/web/middleware/identity"
	"superstartPro/superstar/web/routes"
)

func main() {
	app := bootstrap.New("Superstar database", "一凡Sir")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8081")
}
