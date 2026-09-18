# Go Web Server

A minimal HTTP server built with Go's standard library (`net/http`), demonstrating routing, form handling, and static file serving without any external dependencies.

## Features

- Serves static HTML files from a `static/` directory
- Handles form submissions via `POST` with `r.ParseForm()`
- Custom route handling with method and path validation
- No third-party dependencies — pure Go standard library

## Project Structure

```
.
├── go.mod
├── main.go
├── README.md
└── static
    ├── form.html
    └── index.html
```

## Requirements

- [Go](https://go.dev/dl/) 1.20 or later

## Getting Started

Clone the repository:

```bash
git clone https://github.com/<your-username>/web-server.git
cd web-server
```

Run the server:

```bash
go run main.go
```

The server will start on **http://localhost:8080**.

## Routes

| Route         | Method | Description                                              |
|---------------|--------|------------------------------------------------------------|
| `/`           | GET    | Serves static files from `static/` (defaults to `index.html`) |
| `/form.html`  | GET    | Static page with a name/address form                    |
| `/form`       | POST   | Handles the form submission and echoes back the values   |
| `/hello`      | GET    | Returns a simple `hello!` text response                  |

## Example

Submitting the form at `/form.html` sends a `POST` request to `/form`, which responds with:

```
Post request successful
Name = John Doe
Address = 123 Main St
```

## Roadmap / Ideas

- [ ] Add input validation and HTML escaping for form values
- [ ] Add unit tests for handlers
- [ ] Add a `Makefile` or Taskfile for common commands
- [ ] Containerize with Docker

## License

This project is licensed under the [MIT License](LICENSE).