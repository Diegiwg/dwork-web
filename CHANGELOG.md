# DWork Web Framework

## Current Version: 0.0.10

## Changelog

### Version 0.0.10

    - Modified the Router API to include the Enable (EnableRouter) and RegisterRoute functions as methods of the router object..
    - Now, you first use the router object and utilize its functions as methods to define routes. 

### Version 0.0.9

    - Developed a route collision checking system, which displays error messages using the `log.Error` level and returns an error that developers can handle.
    - Checked for the following situations:
        - Identical static routes.
        - Identical dynamic routes.
        - Collision of parameters in routes.
        - Repeated parameter in a route.

### Version 0.0.8

    - Moved the function for loading custom HTTP handlers from `parse.go` to `router.go`.
    - Added a test suite for the parser functionality.

### Version 0.0.7

    - Added TODOs for route collision checking.
    - Added a test suite for the route registration function.

### Version 0.0.6

    - Modified the route API.
    - Unified the annotation of Handler functions, which now receive a Context object, providing access to the Request, Response, and Params objects.
    - Unified the Route Registration functions.
    - Optimized the parser.
    - The documentation is now marked as a work in progress, allowing for extensive changes to the API.

### Version 0.0.5

    - Improved the documentation.

### Version 0.0.4

    - Improved dynamic route support.
    - Now, it is possible to have dynamic routes, which are not necessarily limited to having the "parameter" as the last part of the route, e.g.
        - RegisterDynamicRoute(..., "/project/:id/name", handler), is a route that will work for the following cases: "project/20/name" and "project/xxx/name".

### Version 0.0.3

    - Added support for simple dynamic routes with a single parameter at the end of the URL.
    - Static routes can override dynamic routes.
    - For example, by registering a static route ("/route/static") and a dynamic route ("/route/:id"), the static handler will be executed when accessing "/route/static". However, any other value in "/route/:value" will trigger the dynamic handler.

### Version 0.0.2

    - Start of work to enable the use of dynamic routes.

### Version 0.0.1

    - Basic routing functionality implemented.
    - Static and dynamic route handling added.
    - Logging system (dwork_logger) integrated.
    - Simple example application included.