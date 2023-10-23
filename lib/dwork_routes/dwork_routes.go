package dwork_routes

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/Diegiwg/dwork-web/lib/dwork_logger"
)

type Route struct {
	path    string
	Content string
}

func MakeRoute() map[string]Route {
	routes := make(map[string]Route)
	return routes
}

func RegisterRoute(routes *map[string]Route, path string, content string) {
	(*routes)[path] = Route{
		path:    path,
		Content: content,
	}
}

func AutoRegisterRoutes(routes *map[string]Route, indexPageName string) {
	// Procura se tem arquivos html dentro da pasta pages
	_, err := os.Stat("pages")
	if os.IsNotExist(err) {
		dwork_logger.Fatal("Pages directory not found!")
	}

	files, err := os.ReadDir("pages")
	if err != nil {
		dwork_logger.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ".html") {
			continue
		}

		name := strings.TrimSuffix(file.Name(), ".html")
		dwork_logger.Info("Registering route " + name)

		content, err := os.ReadFile(path.Join("pages", file.Name()))
		if err != nil {
			dwork_logger.Fatal(err)
		}

		RegisterRoute(routes, "/"+name, string(content))
		if name == indexPageName {
			RegisterRoute(routes, "/", string(content))
		}

		dwork_logger.Success("Route " + name + " registered!")
	}
}

func EnableHandler(routes *map[string]Route) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		route, ok := (*routes)[r.RequestURI]
		if !ok {
			http.NotFound(w, r)
		}

		fmt.Fprint(w, route.Content)
	})
}
