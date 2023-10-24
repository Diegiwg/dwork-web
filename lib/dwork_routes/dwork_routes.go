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
	Param          string
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
				Param:          "",
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
				Param:          "",
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

		param := ""
		if kind == "special" {
			param = strings.TrimPrefix(part, ":")
			part = "@"
		}

		// Check if is the last part
		if i == len(parts)-1 {
			node[part] = &Route{
				Kind:           kind,
				Path:           part,
				Param:          param,
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
				Param:          param,
				Handler:        nil,
				DynamicHandler: nil,
				Routes:         MakeRouter(),
			}
		}

		node = node[part].Routes

	}
}

func parseRoute(routes *Routes, parts []string, route **Route, params *RouteParams) {
	var node Routes = *routes

	for i, part := range parts {

		// If the last part try to get the route
		if i == len(parts)-1 {
			// Nullifier route
			(*route) = nil

			// Check if exist in node
			if _, ok := node[part]; ok {
				*route = node[part]
				continue
			}

			// Check if in node routes exist a special route
			if _, ok := node["@"]; ok {
				*route = node["@"]

				(*params)[(*route).Param] = part
				node = (*route).Routes
				continue
			}
		}

		// Check if exist in node routes
		if _, ok := node[part]; ok {
			node = node[part].Routes
			continue
		}

		// Check if in node routes exist a special route
		if _, ok := node["@"]; ok {
			*route = node["@"]

			(*params)[(*route).Param] = part
			node = (*route).Routes
			continue
		}
	}

	// for i, part := range parts {

	// 	// If the last part try to get the route
	// 	if i == len(parts)-1 {
	// 		if _, ok := node[part]; !ok {
	// 			*route = nil
	// 			continue
	// 		}

	// 		*route = node[part]
	// 	}

	// 	// Check if part exist in map
	// 	if _, ok := node[part]; !ok {
	// 		continue
	// 	}

	// 	*route = node[part]
	// 	node = (*route).Routes
	// }
}

func EnableRouter(routes *Routes) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		path = strings.TrimLeft(path, "/")
		path = strings.TrimRight(path, "/")

		parts := strings.Split(path, "/")
		var route *Route = nil
		params := make(RouteParams)

		parseRoute(routes, parts, &route, &params)

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
			content := route.DynamicHandler(res, req, params)
			fmt.Fprint(res, content)
			return
		}

	})
}
