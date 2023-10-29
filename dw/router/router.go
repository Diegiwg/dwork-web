package router

import "net/http"

var DEBUG_FLAG bool = false

func MakeRouter() Routes {
	return make(map[string]*Route)
}

// Enable enables the Routes.
//
// It sets up the HTTP handler for the Routes. When a request comes in, it parses
// the route and parameters from the URL path and the request method.
//
// If the route is not found or the corresponding handler is nil, it returns
// a 404 Not Found response.
//
// If the route is found, it creates a DWorkContext with the parsed parameters,
// the response writer and the request, and passes it to the route's handler function.
func (routes *Routes) Enable() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		verb := req.Method
		route, params := parse(routes, req.URL.Path, verb)

		if route == nil || route.Handler == nil {
			http.NotFound(res, req)
			return
		}

		context := DWorkContext{
			Params:   params,
			Response: res,
			Request:  req,
		}

		route.Handler(context)
	})
}

// EnableDebug enables the debug mode.
func (routes *Routes) EnableDebug() { DEBUG_FLAG = true }
