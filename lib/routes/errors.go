package routes

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

type InvalidHttpVerb struct {
}

func (err InvalidHttpVerb) Error() string {
	return "Invalid HTTP Verb"
}
