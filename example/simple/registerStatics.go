package main

import (
	dworkweb "github.com/Diegiwg/dwork-web/dw"
)

func registerStatics(app *dworkweb.App) {

	app.GET("/", func(ctx dworkweb.Context) {
		ctx.Response.Html(html("Home Page!"))
	})

	app.GET("/about", func(ctx dworkweb.Context) {
		ctx.Response.Html(html("About Page!"))
	})

	app.GET("/about/project", func(ctx dworkweb.Context) {
		ctx.Response.Html(html("About Project Page!"))
	})
}
