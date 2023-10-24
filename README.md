# DWork Web

DWork Web is a web framework for Go that's lightweight and free from external dependencies. It's designed to simplify web application development by providing essential features. This framework acts as an extension to Go's native net/http package and includes a flexible logging system (dwork_logger) and a robust routing engine (dwork_routes) for creating both static and dynamic routes.

## Table of Contents

- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
- [Features](#features)
  - [Static Routes](#static-routes)
  - [Dynamic Routes](#dynamic-routes)
  - [Logging (dwork_logger)](#logging-dwork_logger)
- [Example](#example)
- [Usage](#usage)
  - [Router Usage](#router-usage)
  - [Static Routes Usage](#static-routes-usage)
  - [Dynamic Routes Usage](#dynamic-routes-usage)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Installation

To kickstart your DWork Web experience, initiate the installation process by using the following command:

```bash
go get github.com/Diegiwg/dwork-web
```

## Features

### Static Routes

DWork Web offers a streamlined approach to define static routes. Each static route is tightly coupled with a specific URL path and linked to a handler function that's responsible for generating the response.

### Dynamic Routes

Dynamic routes are a standout feature of DWork Web. They enable you to define routes with placeholders for capturing values from the URL, simplifying the handling of variable data. With dynamic routes, you can create flexible and data-driven web applications.

### Logging (dwork_logger)

DWork Web seamlessly integrates a logging system based on log levels, offering insights into the server's activities and aiding in issue diagnosis.

## Example

Explore our exemple project that showcases the capabilities of DWork Web. To access it, navigate to the [example](https://github.com/Diegiwg/dwork-web/tree/master/example) folder and compile the project.

## Usage

### Router Usage

To begin utilizing the routing system, you must first create an instance of the routes object. Here's a basic example:

```go
package main

import (
 "net/http"

 "github.com/Diegiwg/dwork-web/lib/dwork_logger"
 "github.com/Diegiwg/dwork-web/lib/dwork_routes"
)

func main() {
 routes := dwork_routes.MakeRouter()
 dwork_routes.EnableRouter(&routes)

 // Register your routes here

 dwork_logger.Info("Server listening on http://localhost:8080")
 http.ListenAndServe(":8080", nil)
}
```

### Static Routes Usage

To start using static routes, employ the `RegisterRoute` function. Pass a reference to the routes object, the desired path, and the handler function. For example:

```go
dwork_routes.RegisterRoute(&routes, "/", func(w http.ResponseWriter, r *http.Request) string {
  return "<h1>Home Page!</h1>"
 })

dwork_routes.RegisterRoute(&routes, "/about", func(w http.ResponseWriter, r *http.Request) string {
  return "<h1>About Page!</h1>"
 })
```

### Dynamic Routes Usage

For dynamic routes, utilize the `RegisterDynamicRoute` function. Pass a reference to the routes object, the desired path, and the handler function. Here's an example:

```go
dwork_routes.RegisterDynamicRoute(&routes, "/user/:id", func(w http.ResponseWriter, r *http.Request) string {
  return "<h1>UserID:" + params["id"] + "</h1>"
 })

dwork_routes.RegisterDynamicRoute(&routes, "/user/:id/posts", func(w http.ResponseWriter, r *http.Request) string {
  return "<h1>Posts of UserID:" + params["id"] + "</h1>"
})

dwork_routes.RegisterDynamicRoute(&routes, "/user/:id/posts/:name", func(w http.ResponseWriter, r *http.Request) string {
  return "<h1>Post: " + params["name"] + " of UserID:" + params["id"] + "</h1>"
})
```

## Contributing

Contributions are warmly welcomed! Whether you're opening issues or submitting pull requests, your efforts can enhance this project.

## License

This project is licensed under the MIT License. For full details, please refer to the [LICENSE](LICENSE) file.
