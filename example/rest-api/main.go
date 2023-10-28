package main

import (
	"net/http"

	"github.com/Diegiwg/dwork-web/dwlogger"
	"github.com/Diegiwg/dwork-web/dwroutes"
)

func main() {
	router := dwroutes.MakeRouter()
	router.Enable()

	router.EnableDebug()

	RegisterUserRoutes(&router)

	router.Dump()
	dwlogger.Info("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
