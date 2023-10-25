package routes

import (
	"strings"
)

func RegisterRoute(routes *Routes, path string, handler RouteHandler) {

	// strip the slash's
	path = strings.TrimLeft(path, "/")
	path = strings.TrimRight(path, "/")

	parts := strings.Split(path, "/")

	var node Routes = *routes
	for i, part := range parts {

		// Check if is a Special part
		param := ""
		kind := "common"
		if strings.Contains(part, ":") {
			kind = "special"
			param = strings.TrimPrefix(part, ":")
			part = "@"
		}

		// Check if is the last part
		if i == len(parts)-1 {
			// TODO: Check if this part exist, if so, return a err

			node[part] = &Route{
				Kind:    kind,
				Path:    part,
				Param:   param,
				Handler: handler,
				Routes:  MakeRouter(),
			}
			continue
		}

		// TODO: Check if in part is a special, and exist in current node, if so, return a err

		// Check if part not exist in map
		if _, ok := node[part]; !ok {
			node[part] = &Route{
				Kind:    kind,
				Path:    part,
				Param:   param,
				Handler: nil,
				Routes:  MakeRouter(),
			}
		}

		node = node[part].Routes

	}
}
