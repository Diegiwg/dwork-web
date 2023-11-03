package main

import (
	dworkweb "github.com/Diegiwg/dwork-web/dw"
)

func main() {

	app := dworkweb.MakeApp()

	app.GET("/", func(ctx dworkweb.Context) {

		content := `<h1>
            Minha Primeira Pagina com o DWork Web (GO + HTML)
        </h1>`

		ctx.Response.Html(content)
	})

	app.Serve(":8080")
}
