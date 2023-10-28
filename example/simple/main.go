package main

import (
	"fmt"
	"net/http"

	"github.com/Diegiwg/dwork-web/lib/logger"
	"github.com/Diegiwg/dwork-web/lib/routes"
)

func html(content ...any) string {
	return fmt.Sprint("<h1>", fmt.Sprint(content...), "</h1>")
}

func registerStatics(router *routes.Routes) {

	router.RegisterRoute(routes.GET, "/", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, html("Home: "))
	})

	router.RegisterRoute(routes.GET, "/about", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, html("About: "))
	})

	router.RegisterRoute(routes.GET, "/about/project", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, html("About Project: "))
	})
}

func registerDynamic(router *routes.Routes) {

	router.RegisterRoute(routes.GET, "/user/<int:id>", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, html("User: (UserID: ", dc.Params["id"], ")"))
	})

	router.RegisterRoute(routes.GET, "/user/<int:id>/project/<string:name>", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, html("Project: (UserID: ", dc.Params["id"], ", ProjectName: ", dc.Params["name"], ")"))
	})

	router.RegisterRoute(routes.GET, "/user/<int:id>/project/<string:name>/edit", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, html("Project Edit: (UserID: ", dc.Params["id"], ", ProjectName: ", dc.Params["name"], ")"))
	})

	router.RegisterRoute(routes.GET, "/user/<int:id>/project/<string:name>/delete", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, html("Project Delete: (UserID: ", dc.Params["id"], ", ProjectName: ", dc.Params["name"], ")"))
	})
}

func registerStaticOverDynamic(router *routes.Routes) {

	router.RegisterRoute(routes.GET, "/user/all", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, html("All Users: "))
	})

	router.RegisterRoute(routes.GET, "/user/<int:id>/project/all", func(dc routes.DWorkContext) {
		fmt.Fprint(dc.Response, html("All Projects of User: (UserID: ", dc.Params["id"], ")"))
	})
}

func registerBadRoutes(router *routes.Routes) {

	str := "Error reported successfully!"

	// PathAlreadyExist
	if err := router.RegisterRoute(routes.GET, "/", nil); err != nil {
		logger.Info(str)
	}

	if err := router.RegisterRoute(routes.GET, "/about", nil); err != nil {
		logger.Info(str)
	}

	if err := router.RegisterRoute(routes.GET, "/user/<int:id>", nil); err != nil {
		logger.Info(str)
	}

	// ParamsConflict
	if err := router.RegisterRoute(routes.GET, "/user/<uuid:id>/all", nil); err != nil {
		logger.Info(str)
	}

	if err := router.RegisterRoute(routes.GET, "/user/<int:id>/project/<uuid:name>", nil); err != nil {
		logger.Info(str)
	}

	// RepeatedParameter
	if err := router.RegisterRoute(routes.GET, "/user/<int:id>/project/<int:id>/edit", nil); err != nil {
		logger.Info(str)
	}

	if err := router.RegisterRoute(routes.GET, "/user/<int:id>/project/<string:name>/link/<uuid:id>", nil); err != nil {
		logger.Info(str)
	}

	// InvalidHttpVerb
	fakeHttpVerb := routes.GET + 100
	if err := router.RegisterRoute(fakeHttpVerb, "/user/<int:id>/project/<string:name>/delete", nil); err != nil {
		logger.Info(str)
	}

	// InvalidParamType
	if err := router.RegisterRoute(routes.GET, "/user/<null:id>/project", nil); err != nil {
		logger.Info(str)
	}

	// InvalidParamStruct
	if err := router.RegisterRoute(routes.GET, "/user/<int:id>/project/string:name", nil); err != nil {
		logger.Info(str)
	}
}

func main() {

	router := routes.MakeRouter()
	router.Enable()

	router.EnableDebug()

	registerStatics(&router)
	registerDynamic(&router)
	registerStaticOverDynamic(&router)
	registerBadRoutes(&router)

	router.Dump()

	logger.Info("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
