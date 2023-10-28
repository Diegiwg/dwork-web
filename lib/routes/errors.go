package routes

import "fmt"

type PathAlreadyExist struct {
	path string
}

func (err PathAlreadyExist) Error() string {
	return fmt.Sprint("The '", err.path, "' path already exist.")
}

type ParamsConflict struct {
	path     string
	param    string
	conflict string
}

func (err ParamsConflict) Error() string {
	return fmt.Sprint("In the '", err.path, "' route there is already a parameter (", err.param, ") in the place where '", err.conflict, "' is being placed.")
}

type RepeatedParameter struct {
	path  string
	param string
}

func (err RepeatedParameter) Error() string {
	return fmt.Sprint("The '", err.param, "' parameter already exists in the '", err.path, "' route.")
}

type InvalidHttpVerb struct{}

func (err InvalidHttpVerb) Error() string {
	return "Not valid HTTP verb in use."
}

type InvalidParamType struct {
	Type  any
	Param string
	Path  string
}

func (err InvalidParamType) Error() string {
	if err.Path == "" {
		err.Path = "index"
	}

	return fmt.Sprint("The '", err.Type, "', in the '", err.Path, "' route is a invalid type for the param '", err.Param, "'.")
}

type InvalidParamStruct struct {
	Param string
	Path  string
}

func (err InvalidParamStruct) Error() string {
	if err.Path == "" {
		err.Path = "index"
	}

	return fmt.Sprint("The '", err.Param, "', in the '", err.Path, "' route is a invalid param struct.")
}
