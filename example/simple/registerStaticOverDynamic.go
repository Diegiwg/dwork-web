package main

import (
	dworkweb "github.com/Diegiwg/dwork-web/dw"
	"github.com/Diegiwg/dwork-web/dw/types"
)

func registerStaticOverDynamic(app *dworkweb.App) {

	app.GET("/user/all", func(ctx dworkweb.Context) {
		ctx.Response.Html(html("User List"))
	})

	app.GET("/user/<int:id>/project/all", func(ctx dworkweb.Context) {

		id, _ := ctx.Request.Params.Int("id")

		ctx.Response.Json(types.Json{
			"action": "all-projects",
			"id":     id,
		})
	})
}
