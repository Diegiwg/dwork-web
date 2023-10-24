# DWork Web Framework

## Current Version: 0.0.5

## Changelog

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
