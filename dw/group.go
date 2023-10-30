package dworkweb

import (
	pathLib "path"
	"strings"

	"github.com/Diegiwg/dwork-web/dw/router"
)

type Group struct {
	app  *App
	path string
}

func (g *Group) GET(path string, handler func(ctx Context)) error {
	return makeRoute(g.app, router.GET, pathLib.Join(g.path, path), handler)
}

func (g *Group) POST(path string, handler func(ctx Context)) error {
	return makeRoute(g.app, router.POST, pathLib.Join(g.path, path), handler)
}

func (g *Group) PUT(path string, handler func(ctx Context)) error {
	return makeRoute(g.app, router.PUT, pathLib.Join(g.path, path), handler)
}

func (g *Group) PATCH(path string, handler func(ctx Context)) error {
	return makeRoute(g.app, router.PATCH, pathLib.Join(g.path, path), handler)
}

func (g *Group) DELETE(path string, handler func(ctx Context)) error {
	return makeRoute(g.app, router.DELETE, pathLib.Join(g.path, path), handler)
}

func (g *Group) HEAD(path string, handler func(ctx Context)) error {
	return makeRoute(g.app, router.HEAD, pathLib.Join(g.path, path), handler)
}

func (g *Group) OPTIONS(path string, handler func(ctx Context)) error {
	return makeRoute(g.app, router.OPTIONS, pathLib.Join(g.path, path), handler)
}

func (app *App) Group(path string) Group {

	// Make all routes for the group
	makeRoute(app, router.GET, path, nil)
	makeRoute(app, router.POST, path, nil)
	makeRoute(app, router.PUT, path, nil)
	makeRoute(app, router.PATCH, path, nil)
	makeRoute(app, router.DELETE, path, nil)
	makeRoute(app, router.HEAD, path, nil)
	makeRoute(app, router.OPTIONS, path, nil)

	return Group{
		app:  app,
		path: strings.TrimSuffix(strings.TrimPrefix(path, "/"), "/"),
	}
}

func (g *Group) Group(path string) Group {

	fullPath := pathLib.Join(g.path, strings.TrimSuffix(strings.TrimPrefix(path, "/"), "/"))

	// Make all routes for the group
	makeRoute(g.app, router.GET, fullPath, nil)
	makeRoute(g.app, router.POST, fullPath, nil)
	makeRoute(g.app, router.PUT, fullPath, nil)
	makeRoute(g.app, router.PATCH, fullPath, nil)
	makeRoute(g.app, router.DELETE, fullPath, nil)
	makeRoute(g.app, router.HEAD, fullPath, nil)
	makeRoute(g.app, router.OPTIONS, fullPath, nil)

	return Group{
		app:  g.app,
		path: fullPath,
	}
}
