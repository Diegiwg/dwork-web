package routes

import (
	"strings"

	"github.com/Diegiwg/dwork-web/lib/logger"
)

type PathAlreadyExist struct {
	path string
}

func (err PathAlreadyExist) Error() string {
	return "Path already exist: '" + err.path + "'"
}

type ParamsConflictInDynamicRoute struct {
	path     string
	param    string
	conflict string
}

func (err ParamsConflictInDynamicRoute) Error() string {
	return "Params conflict in dynamic route: '" + err.path + "'. The '" + err.param + "' parameter already exists, and an attempt was made to add the '" + err.conflict + "' parameter."
}

type SameParamAlreadyExistsInDynamicRoute struct {
	path  string
	param string
}

func (err SameParamAlreadyExistsInDynamicRoute) Error() string {
	return "Same param already exists in dynamic route: '" + err.path + "'. The '" + err.param + "' parameter already exists."
}

func (routes *Routes) RegisterRoute(path string, handler RouteHandler) error {

	parts := strings.Split(strings.TrimLeft(strings.TrimRight(path, "/"), "/"), "/")
	params := make(map[string]bool)

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

			// Check if this part exist, if so, return a err
			if ok := (node)[part]; ok != nil {
				err := PathAlreadyExist{path}
				logger.Error(err)
				return err
			}

			node[part] = &Route{
				Kind:    kind,
				Path:    part,
				Param:   param,
				Handler: handler,
				Routes:  MakeRouter(),
			}
			continue
		}

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

		// Check if exist a conflite in Special routes
		if kind == "special" && node["@"].Param != param {
			err := ParamsConflictInDynamicRoute{path, node["@"].Param, param}
			logger.Error(err)
			return err
		}

		if kind == "special" {
			// Check if param exist in params list
			if temp := params[param]; temp {
				err := SameParamAlreadyExistsInDynamicRoute{path, param}
				logger.Error(err)
				return err
			}

			params[param] = true
		}

		node = node[part].Routes

	}

	return nil
}
