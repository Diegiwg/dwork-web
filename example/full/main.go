package main

import (
	dworkweb "github.com/Diegiwg/dwork-web/dw"
	"github.com/Diegiwg/dwork-web/dw/logger"
	"github.com/Diegiwg/dwork-web/dw/types"
)

func main() {
	app := dworkweb.MakeApp()

	app.GET("/", func(ctx dworkweb.Context) {
		ctx.Response.Html("<h1>Hello World!</h1>")
	})

	app.GET("/user/<int:id>", func(ctx dworkweb.Context) {

		id, err := ctx.Request.Params.Int("id")
		if err != "" {
			logger.Error(err)

			ctx.Response.Status(types.SC_CE_BadRequest)
			ctx.Response.Json(types.Json{
				"error": err,
			})
			return
		}

		ctx.Response.Json(types.Json{
			"id": id,
		})
	})

	app.Serve(":8080")
}