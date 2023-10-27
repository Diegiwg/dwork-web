package routes

import (
	"strings"
)

func parse(routes *Routes, path string, verb string) (*Route, RouteParams) {
	path = strings.TrimRight(strings.TrimLeft(path, "/"), "/")
	parts := strings.Split(path, "/")

	// Parse the Route
	var route *Route = nil
	var node Routes = (*routes)[verb].Routes
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

	return route, params
}
