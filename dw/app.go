package dworkweb

import (
	"net/http"

	"github.com/Diegiwg/dwork-web/dw/logger"
	"github.com/Diegiwg/dwork-web/dw/router"
)

type App struct {
	router router.Routes
}

type Context struct {
	Response CtxResponse
	Request  CtxRequest
}

// Serve
func (app *App) Serve(port string) {
	app.router.Enable()
	logger.Info("Server listening on http://localhost" + port)
	http.ListenAndServe(port, nil)
}

func MakeApp() App {
	app := App{
		router: router.MakeRouter(),
	}

	return app
}

// Routes returns the Router object of the App.
func (app *App) Routes() *router.Routes {
	return &app.router
}

func makeRoute(app *App, method router.HTTPVerb, path string, handler func(ctx Context)) error {
	if handler == nil {
		return app.router.RegisterRoute(method, path, nil)
	}

	return app.router.RegisterRoute(method, path, func(dc router.DWorkContext) {

		response := CtxResponse{
			Raw: &dc.Response,
		}

		request := CtxRequest{
			Raw: dc.Request,
			Params: CtxRequestParams{
				values: &dc.Params,
			},
		}

		context := Context{
			Response: response,
			Request:  request,
		}

		handler(context)
	})
}

// GET handles HTTP GET requests for the specified path.
func (app *App) GET(path string, handler func(ctx Context)) error {
	return makeRoute(app, router.GET, path, handler)
}

// POST handles HTTP POST requests for the specified path.
func (app *App) POST(path string, handler func(ctx Context)) error {
	return makeRoute(app, router.POST, path, handler)
}

// PUT handles HTTP PUT requests for the specified path.
func (app *App) PUT(path string, handler func(ctx Context)) error {
	return makeRoute(app, router.PUT, path, handler)
}

// PATCH handles HTTP PATCH requests for the specified path.
func (app *App) PATCH(path string, handler func(ctx Context)) error {
	return makeRoute(app, router.PATCH, path, handler)
}

// DELETE handles HTTP DELETE requests for the specified path.
func (app *App) DELETE(path string, handler func(ctx Context)) error {
	return makeRoute(app, router.DELETE, path, handler)
}

// HEAD handles HTTP HEAD requests for the specified path.
func (app *App) HEAD(path string, handler func(ctx Context)) error {
	return makeRoute(app, router.HEAD, path, handler)
}

// OPTIONS handles HTTP OPTIONS requests for the specified path.
func (app *App) OPTIONS(path string, handler func(ctx Context)) error {
	return makeRoute(app, router.OPTIONS, path, handler)
}
