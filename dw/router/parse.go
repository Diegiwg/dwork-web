package router

import (
	"strings"
)

// Parse parses the given path and returns the corresponding Route and RouteParams.
//
// It takes the following parameters:
//
// - routes: a pointer to a Routes object, which contains the routes for different HTTP verbs.
//
// - path: a string representing the path to be parsed.
//
// - verb: a string representing the HTTP verb.
//
// It returns the following:
//
// - route: a pointer to the Route object that matches the given path and verb.
//
// - params: a map of string to interface{}, which contains the parsed parameters from the path.
func parse(routes *Routes, path string, verb string) (*Route, RouteParams) {

	parts := strings.Split(strings.TrimRight(strings.TrimLeft(path, "/"), "/"), "/")

	var route *Route = nil
	params := make(RouteParams)

	var node Routes = (*routes)[verb].Routes
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
				temp := node["@"]

				typedValue, err := ParseParamType(part, temp.ParamType, path)
				if err != nil {
					return nil, nil
				}

				params[temp.Param] = typedValue
				route = temp
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
			temp := node["@"]

			// Convert part::value to paramType::value
			typedValue, err := ParseParamType(part, temp.ParamType, path)
			if err != nil {
				return nil, nil
			}

			params[temp.Param] = typedValue
			node = temp.Routes
			continue
		}

		// Nullifier route
		route = nil
		continue
	}

	return route, params
}
