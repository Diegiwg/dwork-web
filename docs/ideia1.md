# DWork Web

## Overview

DWork Web is a web framework for the Go programming language designed to simplify web application development. It is a lightweight framework with no external dependencies, serving as a wrapper for the original Go `net/http` package. DWork Web includes a level-based logging system (dwork_logger) and a flexible routing system (dwork_routes) for both static and dynamic route creation.

## Table of Contents

- [Usage](#usage)
- [API Description](#api-description)
  - [dwork_routes](#dwork_routes)
  - [dwork_logger](#dwork_logger)

## Usage

To get started with DWork Web, follow these steps:

1. **Installation**:

   First, make sure you have Go installed on your system. Then, you can install DWork Web using `go get`:

   ```sh
   go get github.com/yourusername/dworkweb
   ```

2. **Import**:

   Import DWork Web into your Go project:

   ```go
   import "github.com/yourusername/dworkweb"
   ```

3. **Create a Web Application**:

   You can start building your web application using DWork Web. Here's a simple example to create a basic web server:

   ```go
   package main

   import (
       "github.com/yourusername/dworkweb"
   )

   func main() {
       app := dworkweb.New()
       app.Get("/", func(c *dworkweb.Context) {
           c.HTML(200, "Hello, DWork Web!")
       })
       app.Run(":8080")
   }
   ```

4. **Build and Run**:

   Build and run your application:

   ```sh
   go run yourapp.go
   ```

   Your DWork Web application will be accessible at `http://localhost:8080`.

## API Description

### dwork_routes

The `dwork_routes` package allows you to define and manage routes for your web application. You can create both static and dynamic routes, and handle HTTP methods such as GET, POST, PUT, DELETE, etc.

Here's an example of defining a static route:

```go
app := dworkweb.New()
app.Get("/home", func(c *dworkweb.Context) {
    c.HTML(200, "Welcome to the Home Page")
})
```

For dynamic routes, you can use route parameters:

```go
app.Get("/user/:id", func(c *dworkweb.Context) {
    userID := c.Param("id")
    c.String(200, "User ID: "+userID)
})
```

### dwork_logger

The `dwork_logger` package provides a flexible logging system for your DWork Web application. You can log messages at different levels, including INFO, DEBUG, ERROR, and more. Here's how to use it:

```go
logger := dworkweb.NewLogger()
logger.Info("This is an information message.")
logger.Error("This is an error message.")
```

You can customize log output and destinations according to your application's requirements.

For more detailed information and examples on using DWork Web, please refer to the documentation.

---

Feel free to contribute, report issues, or request new features on the [DWork Web GitHub repository](https://github.com/yourusername/dworkweb). We welcome your feedback and contributions to make DWork Web even better!

Happy coding with DWork Web! 🚀
