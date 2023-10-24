package dwork_routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Diegiwg/dwork-web/lib/dwork_logger"
)

type RouteHandler func(http.ResponseWriter, *http.Request) string
type DynamicRouteHandler func(http.ResponseWriter, *http.Request, RouteParams) string

type RouteParams map[string]string

type Route struct {
	Kind           string
	Path           string
	Params         RouteParams
	Handler        RouteHandler
	DynamicHandler DynamicRouteHandler
	Routes         Routes
}

type Routes map[string]*Route

func MakeRouter() Routes {
	return make(map[string]*Route)
}

func RegisterRoute(routes *Routes, path string, handler RouteHandler) {

	// Check if is not a common route
	if strings.Contains(path, ":") {
		dwork_logger.Fatal("Invalid common route: " + path)
	}

	// strip the slash's
	path = strings.TrimLeft(path, "/")
	path = strings.TrimRight(path, "/")

	parts := strings.Split(path, "/")

	var node Routes = *routes
	for i, part := range parts {

		// Check if is the last part
		if i == len(parts)-1 {
			node[part] = &Route{
				Kind:           "common",
				Path:           part,
				Params:         nil,
				Handler:        handler,
				DynamicHandler: nil,
				Routes:         MakeRouter(),
			}
			continue
		}

		// Check if part not exist in map
		if _, ok := node[part]; !ok {
			node[part] = &Route{
				Kind:           "common",
				Path:           part,
				Params:         nil,
				Handler:        nil,
				DynamicHandler: nil,
				Routes:         MakeRouter(),
			}
		}

		node = node[part].Routes

	}
}

func RegisterDynamicRoute(routes *Routes, path string, handler DynamicRouteHandler) {

	// Check if is not a special route
	if !strings.Contains(path, ":") {
		dwork_logger.Fatal("Invalid special route: " + path)
	}

	// strip the slash's
	path = strings.TrimLeft(path, "/")
	path = strings.TrimRight(path, "/")

	parts := strings.Split(path, "/")

	var node Routes = *routes
	for i, part := range parts {

		kind := "common"
		if strings.Contains(part, ":") {
			kind = "special"
		}

		if kind == "special" {
			part = "@"
		}

		// Check if is the last part
		if i == len(parts)-1 {
			node[part] = &Route{
				Kind:           kind,
				Path:           part,
				Params:         nil,
				Handler:        nil,
				DynamicHandler: handler,
				Routes:         MakeRouter(),
			}
			continue
		}

		// Check if part not exist in map
		if _, ok := node[part]; !ok {
			node[part] = &Route{
				Kind:           kind,
				Path:           part,
				Params:         nil,
				Handler:        nil,
				DynamicHandler: nil,
				Routes:         MakeRouter(),
			}
		}

		node = node[part].Routes

	}
}

func parseRoute(node Routes, parts []string, route **Route) {

	// params := make(RouteParams)

	for i, part := range parts {

		// If the last part try to get the route
		if i == len(parts)-1 {
			if _, ok := node[part]; !ok {
				*route = nil
				continue
			}

			*route = node[part]
		}

		// Check if part exist in map
		if _, ok := node[part]; !ok {
			continue
		}

		*route = node[part]
		node = (*route).Routes
	}
}

func EnableRouter(routes *Routes) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		path = strings.TrimLeft(path, "/")
		path = strings.TrimRight(path, "/")

		parts := strings.Split(path, "/")
		var node Routes = *routes
		var route *Route = nil

		parseRoute(node, parts, &route)

		if route == nil {
			http.NotFound(res, req)
			return
		}

		// Common route
		if route.Kind == "common" && route.Handler == nil {
			http.NotFound(res, req)
			return
		}

		if route.Kind == "common" {
			content := route.Handler(res, req)
			fmt.Fprint(res, content)
			return
		}

		if route.Kind == "special" && route.DynamicHandler == nil {
			http.NotFound(res, req)
			return
		}

		if route.Kind == "special" {
			content := route.DynamicHandler(res, req, route.Params)
			fmt.Fprint(res, content)
			return
		}

	})
}
