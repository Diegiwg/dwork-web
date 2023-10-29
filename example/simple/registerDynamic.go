package main

import (
	dworkweb "github.com/Diegiwg/dwork-web/dw"
	"github.com/Diegiwg/dwork-web/dw/types"
)

func registerDynamic(app *dworkweb.App) {

	app.GET("/user/<int:id>", func(ctx dworkweb.Context) {
		id, _ := ctx.Request.Params.Int("id")

		ctx.Response.Json(types.Json{
			"id": id,
		})
	})

	app.GET("/user/<int:id>/project/<string:name>", func(ctx dworkweb.Context) {
		id, _ := ctx.Request.Params.Int("id")
		name, _ := ctx.Request.Params.String("name")

		ctx.Response.Json(types.Json{
			"id":   id,
			"name": name,
		})
	})

	app.GET("/user/<int:id>/project/<string:name>/edit", func(ctx dworkweb.Context) {
		id, _ := ctx.Request.Params.Int("id")
		name, _ := ctx.Request.Params.String("name")

		ctx.Response.Json(types.Json{
			"action": "edit",
			"id":     id,
			"name":   name,
		})
	})

	app.GET("/user/<int:id>/project/<string:name>/delete", func(ctx dworkweb.Context) {
		id, _ := ctx.Request.Params.Int("id")
		name, _ := ctx.Request.Params.String("name")

		ctx.Response.Json(types.Json{
			"action": "delete",
			"id":     id,
			"name":   name,
		})
	})
}
