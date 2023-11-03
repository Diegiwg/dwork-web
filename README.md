# DWork Web

## Project Overview

DWork Web is an experimental web framework designed for simplicity and flexibility.

## Project Documentation

You can find the documentation for this project in [`English`](https://diegiwg.github.io/dwork-web/) or [`Portuguese`](https://diegiwg.github.io/dwork-web/pt).

## Creating Your First Website

To create your first website using DWork Web, follow these steps:

1. **Installation**: Integrate DWork into your project using the following command in the terminal:

   ```bash
   go get github.com/Diegiwg/dwork-web
   ```

2. **Import DWork Web Package**: Import the DWork Web package into your Go project:

   ```go
   package main

   import (
       dworkweb "github.com/Diegiwg/dwork-web/dw"
   )
   ```

3. **Create App Object**: In the `main` function, create an `app` object using the `MakeApp` method:

   ```go
   func main() {
       app := dworkweb.MakeApp()
   }
   ```

4. **Add Routes**: Add routes to your website using methods like `GET`, `POST`, `PUT`, or `DELETE` on the `app` object. For example, to set up a GET route at `/`:

   ```go
   app.GET("/", func(ctx dworkweb.Context) {
       content := `<h1>Your First Page with DWork Web (GO + HTML)</h1>`
       ctx.Response.Html(content)
   })
   ```

5. **Start Server**: Start the server using the `Serve` method:

   ```go
   app.Serve(":8080")
   ```

Now you can access [http://localhost:8080/](http://localhost:8080/) and view your first page.

For a complete example, refer to the [`basic-site`](https://github.com/Diegiwg/dwork-web/tree/master/example/basic-site) directory.

## Examples of Use

You can find more examples of how to use DWork Web in the [`example`](https://github.com/Diegiwg/dwork-web/tree/master/example) directory.

## Contributing

Contributions to DWork Web are highly encouraged and greatly appreciated. Whether you wish to report issues, suggest improvements, or submit pull requests, your contributions are invaluable in enhancing this project.

## Changelog

For details about the latest changes, updates, and version history, please refer to the [`changelog`](CHANGELOG.md) file.

## License

This project is licensed under the MIT License. For comprehensive details, please review the [`license`](LICENSE) file.
