package dwroutes

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/Diegiwg/dwork-web/dwlogger"
)

type RouteParams map[string]any

type DWorkContext struct {
	Params   RouteParams
	Response http.ResponseWriter
	Request  *http.Request
}

type RouteHandler func(DWorkContext)

type Route struct {
	Kind      string
	Path      string
	Param     string
	ParamType ParamTypes
	Handler   RouteHandler
	Routes    Routes
}

type Routes map[string]*Route

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
		dwlogger.Error(err)
		return "", err
	}
	return verbs[verb], nil
}

type ParamTypes string

const (
	EMPTY  ParamTypes = ""
	NULL   ParamTypes = "NULL"
	STRING ParamTypes = "string"
	INT    ParamTypes = "int"
	FLOAT  ParamTypes = "float"
	BOOL   ParamTypes = "bool"
	UUID   ParamTypes = "uuid"
)

func StringToParamType(value string) ParamTypes {
	switch value {

	case "string":
		return STRING
	case "int":
		return INT
	case "float":
		return FLOAT
	case "bool":
		return BOOL
	case "uuid":
		return UUID
	}

	return NULL
}

func ParseParamType(value string, paramType ParamTypes, path string) (interface{}, error) {

	if value == "" || paramType == NULL {
		return nil, InvalidParamType{Type: paramType, Param: value, Path: path}
	}

	switch paramType {

	case STRING:
		return value, nil

	case INT:
		temp, err := strconv.Atoi(value)

		if err != nil {
			return nil, InvalidParamType{Type: paramType, Param: value, Path: path}
		}

		return temp, nil

	case FLOAT:
		temp, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return nil, InvalidParamType{Type: paramType, Param: value, Path: path}
		}

		return temp, nil

	case BOOL:
		temp, err := strconv.ParseBool(value)

		if err != nil {
			return nil, InvalidParamType{Type: paramType, Param: value, Path: path}
		}

		return temp, nil

	case UUID:
		// * Regex from https://github.com/uuidjs/uuid/blob/bc46e198ab06311a9d82d3c9c6222062dd27f760/src/regex.js
		test := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}|00000000-0000-0000-0000-000000000000$`)

		if !test.MatchString(value) {
			return nil, InvalidParamType{Type: paramType, Param: value, Path: path}
		}

		return value, nil

	}

	return nil, InvalidParamType{Type: paramType, Param: value, Path: path}

}
