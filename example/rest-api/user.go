package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Diegiwg/dwork-web/dwroutes"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make(map[int]*User)
var counter = 0

func RegisterUserRoutes(router *dwroutes.Routes) {

	// Create User
	router.RegisterRoute(dwroutes.POST, "/user", func(dc dwroutes.DWorkContext) {

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
	router.RegisterRoute(dwroutes.DELETE, "/user/<int:id>", func(dc dwroutes.DWorkContext) {

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
	router.RegisterRoute(dwroutes.GET, "/user/<int:id>", func(dc dwroutes.DWorkContext) {

		id := dc.Params["id"].(int)

		user, ok := users[id]
		if !ok {
			http.Error(dc.Response, "User not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(dc.Response).Encode(user)
	})
}
