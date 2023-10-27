package main

import (
	"net/http"

	"github.com/Diegiwg/dwork-web/lib/logger"
	"github.com/Diegiwg/dwork-web/lib/routes"
)

func main() {
	router := routes.MakeRouter()
	router.Enable()

	RegisterUserRoutes(&router)

	router.Dump()
	logger.Info("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
