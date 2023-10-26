package main

import (
	"fmt"
	"net/http"

	"github.com/Diegiwg/dwork-web/lib/logger"
	"github.com/Diegiwg/dwork-web/lib/routes"
)

func main() {

	router := routes.MakeRouter()
	router.Enable()

	// * Static Routes

	router.RegisterRoute("/", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Home Page!")
	})
	// ! Register Same Route for Testing purposes
	router.RegisterRoute("/", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Home Page!")
	})

	router.RegisterRoute("/about", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "About Page!")
	})
	router.RegisterRoute("/faq", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "FAQ Page!")
	})

	router.RegisterRoute("/faq/project", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project FAQ Page!")
	})
	router.RegisterRoute("/faq/project", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project FAQ Page!")
	})

	router.RegisterRoute("/project/add", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project Add Page!")
	})

	// ! Dynamic Routes

	router.RegisterRoute("/project/:id", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project ID: "+dc.Params["id"])
	})
	router.RegisterRoute("/project/:id", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project ID: "+dc.Params["id"])
	})

	router.RegisterRoute("/project/:id/name", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project Name: "+dc.Params["id"])
	})
	router.RegisterRoute("/project/:id/name", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project Name: "+dc.Params["id"])
	})
	router.RegisterRoute("/project/:other/x", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project Name: "+dc.Params["id"])
	})

	router.RegisterRoute("/project/:id/:name/show", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project ID: "+dc.Params["id"]+"\nProject Name: "+dc.Params["name"])

	})

	router.RegisterRoute("/project/:id/name/:id/a", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, "Project Name: "+dc.Params["id"])
	})

	// * Server

	logger.Info("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
