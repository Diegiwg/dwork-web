package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Diegiwg/dwork-web/lib/routes"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make(map[int]*User)
var counter = 0

func RegisterUserRoutes(router *routes.Routes) {

	// Create User
	router.RegisterRoute(routes.POST, "/user", func(dc routes.DWorkContext) {

		body, err := io.ReadAll(dc.Request.Body)
		if err != nil {
			http.Error(dc.Response, err.Error(), http.StatusInternalServerError)
			return
		}

		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(dc.Response, err.Error(), http.StatusInternalServerError)
			return
		}

		counter++
		user.Id = counter
		users[user.Id] = &user

		json.NewEncoder(dc.Response).Encode(user)
	})

	// Delete User
	router.RegisterRoute(routes.DELETE, "/user/<int:id>", func(dc routes.DWorkContext) {

		id := dc.Params["id"].(int)

		_, ok := users[id]
		if !ok {
			http.Error(dc.Response, "User not found", http.StatusNotFound)
			return
		}

		delete(users, id)
		fmt.Fprint(dc.Response, "User deleted")
	})

	// Get User
	router.RegisterRoute(routes.GET, "/user/<int:id>", func(dc routes.DWorkContext) {

		id := dc.Params["id"].(int)

		user, ok := users[id]
		if !ok {
			http.Error(dc.Response, "User not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(dc.Response).Encode(user)
	})
}
