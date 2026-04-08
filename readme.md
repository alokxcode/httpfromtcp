# httpfromtcp

HTTP server library built from scratch in Go, on top of raw TCP — no net/http under the hood.

---

## What you can build

- **REST APIs** — full CRUD with dynamic routes
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
