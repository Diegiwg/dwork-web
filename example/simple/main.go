package main

import (
	"fmt"
	"net/http"

	"github.com/Diegiwg/dwork-web/dwlogger"
	"github.com/Diegiwg/dwork-web/dwroutes"
)

func html(content ...any) string {
	return fmt.Sprint("<h1>", fmt.Sprint(content...), "</h1>")
}

func registerStatics(router *dwroutes.Routes) {

	router.RegisterRoute(dwroutes.GET, "/", func(dc dwroutes.DWorkContext) {
		fmt.Fprint(dc.Response, html("Home: "))
	})

	router.RegisterRoute(dwroutes.GET, "/about", func(dc dwroutes.DWorkContext) {
		fmt.Fprint(dc.Response, html("About: "))
	})

	router.RegisterRoute(dwroutes.GET, "/about/project", func(dc dwroutes.DWorkContext) {
		fmt.Fprint(dc.Response, html("About Project: "))
	})
}

func registerDynamic(router *dwroutes.Routes) {

	router.RegisterRoute(dwroutes.GET, "/user/<int:id>", func(dc dwroutes.DWorkContext) {
		fmt.Fprint(dc.Response, html("User: (UserID: ", dc.Params["id"], ")"))
	})

	router.RegisterRoute(dwroutes.GET, "/user/<int:id>/project/<string:name>", func(dc dwroutes.DWorkContext) {
		fmt.Fprint(dc.Response, html("Project: (UserID: ", dc.Params["id"], ", ProjectName: ", dc.Params["name"], ")"))
	})

	router.RegisterRoute(dwroutes.GET, "/user/<int:id>/project/<string:name>/edit", func(dc dwroutes.DWorkContext) {
		fmt.Fprint(dc.Response, html("Project Edit: (UserID: ", dc.Params["id"], ", ProjectName: ", dc.Params["name"], ")"))
	})

	router.RegisterRoute(dwroutes.GET, "/user/<int:id>/project/<string:name>/delete", func(dc dwroutes.DWorkContext) {
		fmt.Fprint(dc.Response, html("Project Delete: (UserID: ", dc.Params["id"], ", ProjectName: ", dc.Params["name"], ")"))
	})
}

func registerStaticOverDynamic(router *dwroutes.Routes) {

	router.RegisterRoute(dwroutes.GET, "/user/all", func(dc dwroutes.DWorkContext) {
		fmt.Fprint(dc.Response, html("All Users: "))
	})

	router.RegisterRoute(dwroutes.GET, "/user/<int:id>/project/all", func(dc dwroutes.DWorkContext) {
		fmt.Fprint(dc.Response, html("All Projects of User: (UserID: ", dc.Params["id"], ")"))
	})
}

func registerBadRoutes(router *dwroutes.Routes) {

	str := "Error reported successfully!"

	// PathAlreadyExist
	if err := router.RegisterRoute(dwroutes.GET, "/", nil); err != nil {
		dwlogger.Info(str)
	}

	if err := router.RegisterRoute(dwroutes.GET, "/about", nil); err != nil {
		dwlogger.Info(str)
	}

	if err := router.RegisterRoute(dwroutes.GET, "/user/<int:id>", nil); err != nil {
		dwlogger.Info(str)
	}

	// ParamsConflict
	if err := router.RegisterRoute(dwroutes.GET, "/user/<uuid:id>/all", nil); err != nil {
		dwlogger.Info(str)
	}

	if err := router.RegisterRoute(dwroutes.GET, "/user/<int:id>/project/<uuid:name>", nil); err != nil {
		dwlogger.Info(str)
	}

	// RepeatedParameter
	if err := router.RegisterRoute(dwroutes.GET, "/user/<int:id>/project/<int:id>/edit", nil); err != nil {
		dwlogger.Info(str)
	}

	if err := router.RegisterRoute(dwroutes.GET, "/user/<int:id>/project/<string:name>/link/<uuid:id>", nil); err != nil {
		dwlogger.Info(str)
	}

	// InvalidHttpVerb
	fakeHttpVerb := dwroutes.GET + 100
	if err := router.RegisterRoute(fakeHttpVerb, "/user/<int:id>/project/<string:name>/delete", nil); err != nil {
		dwlogger.Info(str)
	}

	// InvalidParamType
	if err := router.RegisterRoute(dwroutes.GET, "/user/<null:id>/project", nil); err != nil {
		dwlogger.Info(str)
	}

	// InvalidParamStruct
	if err := router.RegisterRoute(dwroutes.GET, "/user/<int:id>/project/string:name", nil); err != nil {
		dwlogger.Info(str)
	}
}

func main() {

	router := dwroutes.MakeRouter()
	router.Enable()

	router.EnableDebug()

	registerStatics(&router)
	registerDynamic(&router)
	registerStaticOverDynamic(&router)
	registerBadRoutes(&router)

	router.Dump()

	dwlogger.Info("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
