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
