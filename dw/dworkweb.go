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
	logger.Info("Server listening on http://localhost:" + port)
	http.ListenAndServe(port, nil)
}

func MakeApp() App {
	app := App{
		router: router.MakeRouter(),
	}

	return app
}

func (app *App) Routes() *router.Routes {
	return &app.router
}

func makeRoute(app *App, method router.HTTPVerb, path string, handler func(ctx Context)) error {
	return app.router.RegisterRoute(method, path, func(dc router.DWorkContext) {

		response := CtxResponse{
			res: &dc.Response,
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

func (app *App) GET(path string, handler func(ctx Context)) error {
	return makeRoute(app, router.GET, path, handler)
}

func (app *App) POST(path string, handler func(ctx Context)) {
	makeRoute(app, router.POST, path, handler)
}

func (app *App) PUT(path string, handler func(ctx Context)) {
	makeRoute(app, router.PUT, path, handler)
}

func (app *App) PATCH(path string, handler func(ctx Context)) {
	makeRoute(app, router.PATCH, path, handler)
}

func (app *App) DELETE(path string, handler func(ctx Context)) {
	makeRoute(app, router.DELETE, path, handler)
}

func (app *App) HEAD(path string, handler func(ctx Context)) {
	makeRoute(app, router.HEAD, path, handler)
}

func (app *App) OPTIONS(path string, handler func(ctx Context)) {
	makeRoute(app, router.OPTIONS, path, handler)
}
