package dwork_routes

import (
	"fmt"
	"net/http"
	"strings"
)

type RouteHandler func(http.ResponseWriter, *http.Request) string
type DynamicRouteHandler func(http.ResponseWriter, *http.Request, RouteParams) string

type RouteParams map[string]string

type Route struct {
	Kind    string
	Path    string
	Params  RouteParams
	Handler RouteHandler
	Routes  Routes
}

type Routes map[string]*Route

func MakeRouter() Routes {
	return make(map[string]*Route)
}

// Build common route map
func commonRoute(routes *Routes, path string, handler RouteHandler) {
	// strip the slash's
	path = strings.TrimLeft(path, "/")
	path = strings.TrimRight(path, "/")

	parts := strings.Split(path, "/")

	var node Routes = *routes
	for i, part := range parts {

		// Check if is the last part
		if i == len(parts)-1 {
			node[part] = &Route{
				Kind:    "common",
				Path:    part,
				Params:  nil,
				Handler: handler,
				Routes:  MakeRouter(),
			}
			continue
		}

		// Check if part not exist in map
		_, ok := node[part]

		if !ok {
			node[part] = &Route{
				Kind:    "common",
				Path:    part,
				Params:  nil,
				Handler: handler,
				Routes:  MakeRouter(),
			}

			node = node[part].Routes
		}

		if ok {
			node = node[part].Routes
		}

	}
}

// TODO: specialRoute
// * func specialRoute(routes *Routes, path string, handler RouteHandler) {}

func RegisterRoute(routes *Routes, path string, handler RouteHandler) {

	// Check if is not a special route
	if !strings.Contains(path, ":") {
		commonRoute(routes, path, handler)
		return
	}
}

// TODO: RegisterDynamicRoute
func RegisterDynamicRoute(routes *Routes, path string, handler DynamicRouteHandler) {

}

func EnableRouter(routes *Routes) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		path = strings.TrimLeft(path, "/")
		path = strings.TrimRight(path, "/")

		parts := strings.Split(path, "/")
		var node Routes = *routes

		var route *Route = nil
		for i, part := range parts {

			// If the last part try to get the route
			if i == len(parts)-1 {
				if _, ok := node[part]; !ok {
					route = nil
					continue
				}

				route = node[part]
			}

			// Check if part exist in map
			if _, ok := node[part]; !ok {
				continue
			}

			route = node[part]
			node = route.Routes
		}

		if route == nil || route.Handler == nil {
			http.NotFound(res, req)
			return
		}

		content := route.Handler(res, req)
		fmt.Fprint(res, content)
	})
}
