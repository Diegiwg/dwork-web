# 🚀 DWork Web - Documentation

Welcome to the DWork Web documentation, an innovative experimental web framework developed with an emphasis on simplicity and flexibility.

## Installation

To integrate DWork into your project, use the following command in the terminal:

```bash
go get github.com/Diegiwg/dwork-web
```

## Usage (Basic Example)

To start your journey with DWork and create a basic website, begin by importing the package into your Go project:

```go
package main

import (
    dworkweb "github.com/Diegiwg/dwork-web/dw"
)
```

Next, in the `main` function, create an `app` object using the `MakeApp` method:

```go
func main() {
    app := dworkweb.MakeApp()
}
```

To add new routes, select the method corresponding to the desired HTTP verb (`GET`, `POST`, `PUT`, `DELETE`), available on the `app` object. For example, to set up a GET route at `/`:

```go
app.GET("/", func(ctx dworkweb.Context) {
    content := `<h1>My First Page with DWork Web (GO + HTML)</h1>`
    ctx.Response.Html(content)
})
```

To start the server, use the `Serve` method:

```go
app.Serve(":8080")
```

Now you can access [http://localhost:8080/](http://localhost:8080/) and view your first page.

- The complete code for this example is available in the [`basic-site`](https://github.com/Diegiwg/dwork-web/tree/master/example/basic-site) directory.

## Next Steps

Explore the documentation for the `App` module at [`App Module`](https://diegiwg.github.io/dwork-web/en/modules/app) to understand how to expand your first website. Learn how to add more routes, dynamic routes, and other HTTP verbs to develop more complex and interactive web applications. This knowledge will be a valuable addition to your resume, demonstrating your ability to build robust web applications with DWork Web. Begin this journey to success with DWork Web! 🚀
