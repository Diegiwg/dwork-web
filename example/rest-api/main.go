package main

import (
	dworkweb "github.com/Diegiwg/dwork-web/dw"
)

func main() {
	app := dworkweb.MakeApp()
	app.Routes().EnableDebug()

	RegisterUserRoutes(&app)

	app.Serve(":8080")
}
