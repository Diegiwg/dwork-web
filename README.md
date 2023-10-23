# DWork Web

DWork Web is a lightweight and dependency-free web framework for Go, designed to simplify web application development. It acts as a wrapper for the original `net/http` package in Go, enhancing it with additional features. The framework includes a flexible logging system (dwork_logger) and a powerful routing engine (dwork_routes) that enables the creation of both static and dynamic routes.

## Table of Contents

- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
- [Features](#features)
  - [Static Routes](#static-routes)
  - [Dynamic Routes](#dynamic-routes)
  - [Logging (dwork_logger)](#logging-dwork_logger)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Installation

To install DWork Web, you can use `go get`:

```bash
go get github.com/Diegiwg/dwork-web
```

### Usage

Here's an example of how to use DWork Web in your Go application:

```go
package main

import (
    "net/http"
    "github.com/Diegiwg/dwork-web/lib/dwork_logger"
    "github.com/Diegiwg/dwork-web/lib/dwork_routes"
)

func main() {
    // Create a new router
    routes := dwork_routes.MakeRouter()
    dwork_routes.EnableRouter(&routes)

    // Static Routes
    dwork_routes.RegisterRoute(&routes, "/", func(w http.ResponseWriter, r *http.Request) string {
        return "Home Page!"
    })
    dwork_routes.RegisterRoute(&routes, "/about", func(w http.ResponseWriter, r *http.Request) string {
        return "About Page!"
    })
    dwork_routes.RegisterRoute(&routes, "/faq", func(w http.ResponseWriter, r *http.Request) string {
        return "FAQ Page!"
    })
    dwork_routes.RegisterRoute(&routes, "/faq/project", func(w http.ResponseWriter, r *http.Request) string {
        return "Project FAQ Page!"
    })

    // Dynamic Routes
    dwork_routes.RegisterDynamicRoute(&routes, "/project/:id", func(w http.ResponseWriter, r *http.Request, params dwork_routes.RouteParams) string {
        return "Project ID: " + params["id"]
    })
    dwork_routes.RegisterDynamicRoute(&routes, "/project/:id/:name", func(w http.ResponseWriter, r *http.Request, params dwork_routes.RouteParams) string {
        return "Project ID: " + params["id"] + "\nProject Name: " + params["name"]
    })

    // Server
    dwork_logger.Info("Server listening on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

## Features

### Static Routes

DWork Web allows you to define static routes easily. Each static route maps to a specific URL path and is associated with a handler function that generates the response.

### Dynamic Routes

Dynamic routes are a powerful feature of DWork Web. You can define routes with placeholders that capture values from the URL, making it easy to work with variable data. These dynamic routes enable you to create flexible and data-driven web applications.

### Logging (dwork_logger)

DWork Web includes a logging system based on log levels, which helps you keep track of the server's activities and troubleshoot any issues.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to help improve this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
