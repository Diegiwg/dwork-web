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

func recursiveRegisterRoutes(routes *map[string]Route, indexPageName string, dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		dwork_logger.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			recursiveRegisterRoutes(routes, indexPageName, path.Join(dir, file.Name()))
			continue
		}

		if !strings.HasSuffix(file.Name(), ".html") {
			continue
		}

		name := strings.TrimSuffix(file.Name(), ".html")
		content, err := os.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			dwork_logger.Fatal(err)
		}

		path := "/"
		if dir != "pages" {
			path = strings.Split(dir, "pages")[1] + "/"
		}

		dwork_logger.Info("Registering route " + path + name)

		RegisterRoute(routes, path+name, string(content))
		if name == indexPageName {
			RegisterRoute(routes, "/", string(content))
		}

		dwork_logger.Success("Route " + name + " registered!")
	}
}

func AutoRegisterRoutes(routes *map[string]Route, indexPageName string) {
	// Procura se tem arquivos html dentro da pasta pages
	_, err := os.Stat("pages")
	if os.IsNotExist(err) {
		dwork_logger.Fatal("Pages directory not found!")
	}

	recursiveRegisterRoutes(routes, indexPageName, "pages")
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
