package dworkweb

import (
	"encoding/json"
	"fmt"
	"net/http"

	types "github.com/Diegiwg/dwork-web/dw/types"
)

type CtxResponse struct {
	Raw *http.ResponseWriter
}

func (ctx *CtxResponse) Status(code types.StatusCode) {
	(*ctx.Raw).WriteHeader(code.Parse())
}

func (ctx *CtxResponse) Json(data interface{}) {
	res := *ctx.Raw

	// Set content type
	res.Header().Set("Content-Type", "application/json")

	// Send Json
	json.NewEncoder(res).Encode(data)
}

func (ctx *CtxResponse) Text(data string) {
	res := *ctx.Raw

	// Set content type
	res.Header().Set("Content-Type", "text/plain")

	// Send Data
	fmt.Fprint(res, data)
}

func (ctx *CtxResponse) Html(data string) {
	res := *ctx.Raw

	// Set content type
	res.Header().Set("Content-Type", "text/html")

	// Send Data
	fmt.Fprint(res, data)
}
