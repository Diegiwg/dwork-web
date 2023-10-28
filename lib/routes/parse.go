package routes

import (
	"strings"
)

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
