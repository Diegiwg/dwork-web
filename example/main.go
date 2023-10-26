package main

import (
	"fmt"
	"net/http"

	"github.com/Diegiwg/dwork-web/lib/dwork_logger"
	"github.com/Diegiwg/dwork-web/lib/routes"
)

func main() {

	r := routes.MakeRouter()
	routes.EnableRouter(&r)

	// * Static Routes

	routes.RegisterRoute(&r, "/", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Home Page!")
	})
	// ! Register Same Route for Testing purposes
	routes.RegisterRoute(&r, "/", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Home Page!")
	})

	routes.RegisterRoute(&r, "/about", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "About Page!")
	})
	routes.RegisterRoute(&r, "/faq", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "FAQ Page!")
	})

	routes.RegisterRoute(&r, "/faq/project", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project FAQ Page!")
	})
	routes.RegisterRoute(&r, "/faq/project", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project FAQ Page!")
	})

	routes.RegisterRoute(&r, "/project/add", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project Add Page!")
	})

	// ! Dynamic Routes

	routes.RegisterRoute(&r, "/project/:id", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project ID: "+dc.Params["id"])
	})
	routes.RegisterRoute(&r, "/project/:id", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project ID: "+dc.Params["id"])
	})

	routes.RegisterRoute(&r, "/project/:id/name", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project Name: "+dc.Params["id"])
	})
	routes.RegisterRoute(&r, "/project/:id/name", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project Name: "+dc.Params["id"])
	})
	routes.RegisterRoute(&r, "/project/:other/x", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project Name: "+dc.Params["id"])
	})

	routes.RegisterRoute(&r, "/project/:id/:name/show", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project ID: "+dc.Params["id"]+"\nProject Name: "+dc.Params["name"])

	})

	routes.RegisterRoute(&r, "/project/:id/name/:id/a", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project Name: "+dc.Params["id"])
	})

	// * Server

	dwork_logger.Info("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
