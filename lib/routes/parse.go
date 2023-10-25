package routes

import (
	"net/http"
	"strings"
)

func EnableRouter(routes *Routes) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		path := strings.TrimRight(strings.TrimLeft(req.URL.Path, "/"), "/")
		parts := strings.Split(path, "/")

		// Parse the Route
		var route *Route = nil
		var node Routes = *routes
		params := make(RouteParams)

		for i, part := range parts {

			// If the last part try to get the route
			if i == len(parts)-1 {
				// Check if exist in node
				if _, ok := node[part]; ok {
					route = node[part]
					continue
				}

				// Check if in node routes exist a special route
				if _, ok := node["@"]; ok {
					route = node["@"]
					params[route.Param] = part
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
				r := node["@"]

				params[r.Param] = part
				node = r.Routes
				continue
			}

			// Nullifier route
			route = nil
			continue
		}

		// End of Parser

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
