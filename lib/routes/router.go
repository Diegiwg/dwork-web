package routes

import "net/http"

type RouteParams map[string]string

type DWorkContext struct {
	Params   RouteParams
	Response http.ResponseWriter
	Request  *http.Request
}

type RouteHandler func(DWorkContext)

type Route struct {
	Kind    string
	Path    string
	Param   string
	Handler RouteHandler
	Routes  Routes
}

type Routes map[string]*Route

func MakeRouter() Routes {
	return make(map[string]*Route)
}

func (routes *Routes) Enable() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		verb := req.Method
		route, params := parse(routes, req.URL.Path, verb)

		if route == nil || route.Handler == nil {
			http.NotFound(res, req)
			return
		}

		context := DWorkContext{
			Params:   params,
			Response: res,
			Request:  req,
		}

		route.Handler(context)
	})
}
