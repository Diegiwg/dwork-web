package main

import (
	"net/http"

	"github.com/Diegiwg/dwork-web/lib/dwork_logger"
	"github.com/Diegiwg/dwork-web/lib/dwork_routes"
)

func main() {
	routes := dwork_routes.MakeRouter()
	dwork_routes.EnableRouter(&routes)

	// * Static Routes

	dwork_routes.RegisterRoute(&routes, "/", func(w http.ResponseWriter, r *http.Request) string {
		return "Home Page!"
	})
	dwork_routes.RegisterRoute(&routes, "/about", func(w http.ResponseWriter, r *http.Request) string {
		return "<h1>About Page!</h1>"
	})
	dwork_routes.RegisterRoute(&routes, "/faq", func(w http.ResponseWriter, r *http.Request) string {
		return "FAQ Page!"
	})
	dwork_routes.RegisterRoute(&routes, "/faq/project", func(w http.ResponseWriter, r *http.Request) string {
		return "Project FAQ Page!"
	})

	dwork_routes.RegisterRoute(&routes, "/project/add", func(w http.ResponseWriter, r *http.Request) string {
		return "Project Add Page!"
	})

	// ! Dynamic Routes

	dwork_routes.RegisterDynamicRoute(&routes, "/project/:id", func(w http.ResponseWriter, r *http.Request, params dwork_routes.RouteParams) string {
		return "Project ID: " + params["id"]
	})

	dwork_routes.RegisterDynamicRoute(&routes, "/project/:id/name", func(w http.ResponseWriter, r *http.Request, params dwork_routes.RouteParams) string {
		return "Project ID: " + params["id"]
	})

	// * Server

	dwork_logger.Info("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
