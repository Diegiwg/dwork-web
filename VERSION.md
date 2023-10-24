# DWork Web Framework

## Current Version: 0.0.3

## Changelog

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
