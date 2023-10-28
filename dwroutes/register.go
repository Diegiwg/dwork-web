package dwroutes

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Diegiwg/dwork-web/dwlogger"
)

// RegisterRoute registers a new route in the Routes struct.
//
// The function takes in the HTTPVerb, path, and handler as parameters.
//
// It returns an error if there's an issue parsing the verb or if there
// are conflicts with the path or parameters.
func (routes *Routes) RegisterRoute(verb HTTPVerb, path string, handler RouteHandler) error {

	validVerb, err := verb.Parse()
	if err != nil {
		return err
	}

	// Create the root node for the verb if not exist
	if _, ok := (*routes)[validVerb]; !ok {
		(*routes)[validVerb] = &Route{
			Kind:      "METHOD",
			Path:      "",
			Param:     "",
			ParamType: NULL,
			Handler:   nil,
			Routes:    MakeRouter(),
		}
	}

	parts := strings.Split(strings.TrimLeft(strings.TrimRight(path, "/"), "/"), "/")
	params := make(map[string]ParamTypes)

	var node Routes = (*routes)[validVerb].Routes
	for i, part := range parts {

		param := ""
		kind := "common"

		// * Handle the special part's

		var paramType ParamTypes = NULL
		if strings.Contains(part, ":") {

			// Get the type and name of param
			match := (regexp.MustCompile(`<(?P<type>[\w]+):(?P<name>[\w\d]+)>`)).FindStringSubmatch(part)

			if len(match) != 3 {
				err := InvalidParamStruct{Param: part, Path: part}
				dwlogger.Error(err)
				return err
			}

			parsedType := StringToParamType(match[1])
			if parsedType == NULL {
				err := InvalidParamType{Type: match[1], Param: part}
				dwlogger.Error(err)
				return err
			}

			paramType = parsedType
			param = match[2]

			kind = "special"
			part = "@"

		}

		// * Check conflict of equal parameters in the path, and so, returns an error
		if kind == "special" {
			if temp := params[param]; temp != EMPTY {
				err := RepeatedParameter{path, param}
				dwlogger.Error(err)
				return err
			}

			params[param] = paramType
		}

		// * Handle the last part
		if i == len(parts)-1 {

			// * Check for conflict of param in current part of the path, and so, returns an error
			if kind == "special" && node["@"] != nil && (node["@"].Param != param || node["@"].ParamType != paramType) {
				err := ParamsConflict{path, node["@"].Param, param}
				dwlogger.Error(err)
				return err
			}

			// * Check if the part already exist, and if so, returns an error
			if ok := (node)[part]; ok != nil {
				err := PathAlreadyExist{path}
				dwlogger.Error(err)
				return err
			}

			node[part] = &Route{
				Kind:      kind,
				Path:      part,
				Param:     param,
				ParamType: paramType,
				Handler:   handler,
				Routes:    MakeRouter(),
			}
			continue
		}

		// * If the part not exist in node, make it
		if _, ok := node[part]; !ok {
			node[part] = &Route{
				Kind:      kind,
				Path:      part,
				Param:     param,
				ParamType: paramType,
				Handler:   nil,
				Routes:    MakeRouter(),
			}
		}

		// * Check for conflict of param in current part of the path, and so, returns an error
		if kind == "special" && (node["@"].Param != param || node["@"].ParamType != paramType) {
			err := ParamsConflict{path, node["@"].Param, param}
			dwlogger.Error(err)
			return err
		}

		node = node[part].Routes

	}

	if DEBUG_FLAG {
		dwlogger.Debug("Registered route: " + path + " with params: " + fmt.Sprint(params))
	}

	return nil
}
