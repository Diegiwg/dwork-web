package main

import (
	"encoding/json"
	"fmt"
	"io"

	dworkweb "github.com/Diegiwg/dwork-web/dw"
	"github.com/Diegiwg/dwork-web/dw/types"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make(map[int]*User)
var counter = 0

func RegisterUserRoutes(app *dworkweb.App) {

	// Create User
	app.POST("/user", func(ctx dworkweb.Context) {

		body, err := io.ReadAll(ctx.Request.Raw.Body)
		if err != nil {
			ctx.Response.Status(types.SC_CE_BadRequest)
			ctx.Response.Json(types.Json{
				"error": err,
			})
			return
		}

		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			ctx.Response.Status(types.SC_CE_BadRequest)
			ctx.Response.Json(types.Json{
				"error": err,
			})
			return
		}

		counter++
		user.Id = counter
		users[user.Id] = &user

		ctx.Response.Json(user)
	})

	// Delete User
	app.DELETE("/user/<int:id>", func(ctx dworkweb.Context) {

		id, _ := ctx.Request.Params.Int("id")

		_, ok := users[id]
		if !ok {
			ctx.Response.Status(types.SC_CE_NotFound)
			ctx.Response.Json(types.Json{
				"error": "User not Founded",
			})
			return
		}

		delete(users, id)
		ctx.Response.Json(types.Json{
			"message": fmt.Sprintf("User %d deleted", id),
		})
	})

	// Get User
	app.GET("/user/<int:id>", func(ctx dworkweb.Context) {

		id, _ := ctx.Request.Params.Int("id")

		user, ok := users[id]
		if !ok {
			ctx.Response.Status(types.SC_CE_NotFound)
			ctx.Response.Json(types.Json{
				"error": "User not Founded",
			})
			return
		}

		ctx.Response.Json(user)
	})
}
