package main

import (
	"fmt"

	dworkweb "github.com/Diegiwg/dwork-web/dw"
)

func html(content ...any) string {
	return fmt.Sprint("<h1>", fmt.Sprint(content...), "</h1>")
}

func main() {
	app := dworkweb.MakeApp()
	app.Routes().EnableDebug()

	registerStatics(&app)
	registerDynamic(&app)
	registerStaticOverDynamic(&app)
	registerBadRoutes(&app)

	app.Routes().Dump()
	app.Serve(":8080")
}
