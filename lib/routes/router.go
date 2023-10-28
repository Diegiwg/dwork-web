package routes

import "net/http"

var DEBUG_FLAG bool = false

func MakeRouter() Routes {
	return make(map[string]*Route)
}

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

func (routes *Routes) EnableDebug() { DEBUG_FLAG = true }
