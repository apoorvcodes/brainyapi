package main

import (
	"github.com/gominima/cors"
	"github.com/gominima/middlewares"
	"github.com/gominima/minima"
)

func main() {
	app := minima.New()
	crs := cors.New()
	app.UseRaw(middleware.AllowContentEncoding())
	app.UseRaw(middleware.DefaultLogger)
	app.UseRaw(middleware.RequestID)
	app.Use(crs.AllowAll())
        app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.Send("Hello World")
	})
	app.Listen(":3000")
}