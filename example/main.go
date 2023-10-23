package main

import (
	"net/http"

	"github.com/Diegiwg/dwork-web/lib/dwork_logger"
	"github.com/Diegiwg/dwork-web/lib/dwork_routes"
)

func main() {
	routes := dwork_routes.MakeRoute()
	dwork_routes.AutoRegisterRoutes(&routes, "home")
	dwork_routes.EnableHandler(&routes)

	dwork_logger.Info("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
