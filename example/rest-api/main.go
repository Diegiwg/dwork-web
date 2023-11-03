package main

import (
	dworkweb "github.com/Diegiwg/dwork-web/dw"
)

func main() {
	app := dworkweb.MakeApp()
	app.Routes().EnableDebug()

	RegisterUserRoutes(&app)

	app.Routes().Dump()
	app.Serve(":8080")
}
