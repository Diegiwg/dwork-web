package routes

import (
	"strings"

	"github.com/Diegiwg/dwork-web/lib/logger"
)

type HTTPVerb int

const (
	GET HTTPVerb = iota
	POST
	PUT
	PATCH
	DELETE
	HEAD
	OPTIONS
)

func (verb HTTPVerb) Parse() (string, error) {
	verbs := [...]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	if verb < GET || verb > OPTIONS {
		err := InvalidHttpVerb{}
		logger.Error(err)
		return "", err
	}
	return verbs[verb], nil
}

func (routes *Routes) RegisterRoute(verb HTTPVerb, path string, handler RouteHandler) error {

	validVerb, err := verb.Parse()
	if err != nil {
		return err
	}

	// Create the root node for the verb if not exist
	if _, ok := (*routes)[validVerb]; !ok {
		(*routes)[validVerb] = &Route{
			Kind:    "METHOD",
			Path:    "",
			Param:   "",
			Handler: nil,
			Routes:  MakeRouter(),
		}
	}

	parts := strings.Split(strings.TrimLeft(strings.TrimRight(path, "/"), "/"), "/")
	params := make(map[string]bool)

	var node Routes = (*routes)[validVerb].Routes
	for i, part := range parts {

		// * Handle the special part's
		param := ""
		kind := "common"
		if strings.Contains(part, ":") {
			kind = "special"
			param = strings.TrimPrefix(part, ":")
			part = "@"
		}

		// * Handle the last part
		if i == len(parts)-1 {

			// * Check if the part already exist, and if so, returns an error
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

		// * If the part not exist in node, make it
		if _, ok := node[part]; !ok {
			node[part] = &Route{
				Kind:    kind,
				Path:    part,
				Param:   param,
				Handler: nil,
				Routes:  MakeRouter(),
			}
		}

		// * Check for conflict of param in current part of the path, and so, returns an error
		if kind == "special" && node["@"].Param != param {
			err := ParamsConflictInDynamicRoute{path, node["@"].Param, param}
			logger.Error(err)
			return err
		}

		if kind == "special" {
			// * Check conflict of equal parameters in the path, and so, returns an error
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
