package dworkweb

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Diegiwg/dwork-web/dw/router"
)

type ErrorParamNotFound struct {
	Param string
}

func (err ErrorParamNotFound) Error() string {
	return fmt.Sprint("The '", err.Param, "' param not found")
}

type ErrorParamNotIsT struct {
	Param string
	Type  string
}

func (err ErrorParamNotIsT) Error() string {
	return fmt.Sprint("The '", err.Param, "' param is not of type '", err.Type, "'")
}

type CtxRequest struct {
	Raw    *http.Request
	Params CtxRequestParams
}

// Body returns the body data of the Request.
func (ctx *CtxRequest) Body() ([]byte, error) {
	body, err := io.ReadAll(ctx.Raw.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type CtxRequestParams struct {
	values *router.RouteParams
}

// String returns the value of a specified parameter as a string.
//
// Parameters:
//
// - param: the name of the parameter to retrieve.
//
// Return type:
//
// - string: the value of the parameter as a string.
//
// - error: an error of type ErrorParamNotFound if the parameter does not exist.
func (ctx *CtxRequestParams) String(param string) (string, error) {

	// Check of the param exists
	raw, ok := (*ctx.values)[param]
	if !ok {
		return "", ErrorParamNotFound{param}
	}

	return raw.(string), nil
}

// Int checks if the param exists and casts it to an int.
//
// Parameters:
//
// - param: the name of the parameter to retrieve.
//
// Return type:
//
// - int: the value of the parameter as an int.
//
// - error: an error of type ErrorParamNotFound if the parameter does not exist.
func (ctx *CtxRequestParams) Int(param string) (int, error) {

	// Check of the param exists
	raw, ok := (*ctx.values)[param]
	if !ok {
		return 0, ErrorParamNotFound{param}
	}

	// Cast to int
	value, ok := raw.(int)
	if !ok {
		return 0, ErrorParamNotIsT{param, "int"}
	}

	return value, nil
}

// Float checks if the param exists and casts it to a float.
//
// Parameters:
//
// - param: the name of the parameter to retrieve.
//
// Return type:
//
// - float64: the value of the parameter as a float.
//
// - error: an error of type ErrorParamNotFound if the parameter does not exist.
func (ctx *CtxRequestParams) Float(param string) (float64, error) {

	// Check of the param exists
	raw, ok := (*ctx.values)[param]
	if !ok {
		return 0, ErrorParamNotFound{param}
	}

	// Cast to float
	value, ok := raw.(float64)
	if !ok {
		return 0, ErrorParamNotIsT{param, "float"}
	}

	return value, nil
}

// Bool checks if the param exists and casts it to a bool.
//
// Parameters:
//
// - param: the name of the parameter to retrieve.
//
// Return type:
//
// - bool: the value of the parameter as a bool.
//
// - error: an error of type ErrorParamNotFound if the parameter does not exist.
func (ctx *CtxRequestParams) Bool(param string) (bool, error) {

	// Check of the param exists
	raw, ok := (*ctx.values)[param]
	if !ok {
		return false, ErrorParamNotFound{param}
	}

	// Cast to bool
	value, ok := raw.(bool)
	if !ok {
		return false, ErrorParamNotIsT{param, "bool"}
	}

	return value, nil
}

// UUID checks if the param exists and casts it to a uuid.
//
// Parameters:
//
// - param: the name of the parameter to retrieve.
//
// Return type:
//
// - string: the value of the parameter as a uuid.
//
// - error: an error of type ErrorParamNotFound if the parameter does not exist.
func (ctx *CtxRequestParams) UUID(param string) (string, error) {

	// Check of the param exists
	raw, ok := (*ctx.values)[param]
	if !ok {
		return "", ErrorParamNotFound{param}
	}

	// Cast to uuid
	value, ok := raw.(string)
	if !ok {
		return "", ErrorParamNotIsT{param, "uuid"}
	}

	return value, nil
}
