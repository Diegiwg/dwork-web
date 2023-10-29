package dworkweb

import (
	"net/http"

	"github.com/Diegiwg/dwork-web/dw/router"
)

type CtxRequest struct {
	Raw    *http.Request
	Params CtxRequestParams
}

type CtxRequestParams struct {
	values *router.RouteParams
}

func (ctx *CtxRequestParams) Int(param string) (int, string) {

	// Check of the param exists
	raw, ok := (*ctx.values)[param]
	if !ok {
		return 0, "ERROR: Param Not Found"
	}

	// Cast to int
	value, ok := raw.(int)
	if !ok {
		return 0, "ERROR: Param is not an int"
	}

	return value, ""
}

func (ctx *CtxRequestParams) String(param string) (string, string) {

	// Check of the param exists
	raw, ok := (*ctx.values)[param]
	if !ok {
		return "", "ERROR: Param Not Found"
	}

	return raw.(string), ""
}
