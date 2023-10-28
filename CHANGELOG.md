# DWork Web Framework

## Current Version: 0.0.13

## Changelog

### Version 0.0.13

    - Routes::General
        - Error and type structures have been moved to separate files.
        - `Router.Dumper` has been modified to display the 'type' of the parameter.
        - Added the option to enable Debug mode for the Router with Router.EnableDebug. Currently, this Debug mode provides verbose information about registered routes.
    - Routes::Errors
        - Improved error messages with a focus on providing more context to developers.
    - Routes::Typing
        - Route parameters are now "strongly" typed.
        - When defining a parameter in a route, you must provide a type annotation in the format: <type:param>. Supported types include string, int, float, bool, and uuid.
        - Parameters are checked at runtime, and if a user provides a value that is incompatible with the defined parameter type, a generic 404 error is returned.
    - Routes::Tests
        - Tests have been updated to align with the new type system.
        - New tests have been created for type parsing.
    - Examples::General
        - Examples have been updated to match the new type system.
    - TODOS:
        - There is a need to manually cast parameters from the URL to their respective types, e.g., `id := dc.Params["id"].(int)`. The idea is to hide Params as a 'raw' variable and attach methods like parseInt(), parseFloat(), parseBool(), parseUUID(), maybe parseString() for safety and sanitization,  to the `context dc`.
        - Implement descriptive errors for users in case of parameter type mismatches when a 'value' is provided.

### Version 0.0.12

    - Added support for other HTTP verbs.
    - Introduced a new example of a RESTful API that uses GET, POST, and DELETE.
    - Added a Dumper function to print the tree of routes.

### Version 0.0.11

    - Errors have been moved to a separate file for better organization.
    - The implementation of support for HTTP Verbs has been initiated.

### Version 0.0.10

    - Modified the Router API to include the Enable (EnableRouter) and RegisterRoute functions as methods of the router object..
    - Now, you first use the router object and utilize its functions as methods to define routes.
    - The logging module has been moved to lib/logger (formerly lib/dwork_logger).

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
