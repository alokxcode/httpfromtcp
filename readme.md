# httpfromtcp

HTTP server library built from scratch in Go, on top of raw TCP — no net/http under the hood.

---

## What you can build

- **REST APIs** — full CRUD with dynamic routes
- **Static file servers** — read files from disk, serve over HTTP
- **Webhooks** — listen for POST requests from external services
- **Local dev tools** — lightweight server for local tooling or CLIs
- **Learning projects** — understand HTTP at the wire level by reading the source

---

## Install

```bash
go get github.com/alokxcode/httpfromtcp
```

---

## Quickstart

```go
package main

import "github.com/alokxcode/httpfromtcp"

func main() {
    server := httpfromtcp.NewServer()

    server.Handle("GET /", func(req *httpfromtcp.Req, rw *httpfromtcp.ResponseWriter) {
        rw.WriteHeader(httpfromtcp.StatusOk)
        rw.Header().Set("Content-Type", "text/plain")
        rw.Write("Hello, world!")
    })

    server.ListenAndServe(":8080")
}
```

---

## Routing

Register routes using `method + path`:

```go
server.Handle("GET /users", listUsers)
server.Handle("POST /users", createUser)
server.Handle("DELETE /users/{id}", deleteUser)
```

### Dynamic path parameters

Use `{param}` syntax to capture values from the URL:

```go
server.Handle("GET /users/{id}", func(req *httpfromtcp.Req, rw *httpfromtcp.ResponseWriter) {
    id := req.PathValue["id"]
    rw.WriteHeader(httpfromtcp.StatusOk)
    rw.Write("User: " + id.(string))
})
```

---

## Response

### Status codes

```go
rw.WriteHeader(httpfromtcp.StatusOk)                  // 200
rw.WriteHeader(httpfromtcp.StatusCreated)             // 201
rw.WriteHeader(httpfromtcp.StatusBadRequest)          // 400
rw.WriteHeader(httpfromtcp.StatusNotFound)            // 404
rw.WriteHeader(httpfromtcp.StatusInternalServerError) // 500
```

### Headers

```go
rw.Header().Set("Content-Type", "application/json")
rw.Header().Set("X-Request-Id", "abc123")
```

### Body

```go
rw.Write(`{"message": "ok"}`)
```

`Content-Length` is set automatically.

---

## Project structure

```
httpfromtcp/
├── server.go       # Server, Handle(), ListenAndServe()
├── tcp.go          # TCP listener, connection loop
├── handler.go      # Route matching, path param extraction
├── response.go     # ResponseWriter
├── header.go       # Header type
├── status.go       # Status codes
```

---
