package main

import "github.com/alokxcode/httpfromtcp"

func main() {
	router := httpfromtcp.NewServer()

	router.Handle("GET /", func(r httpfromtcp.Req, w *httpfromtcp.ResponseWriter) { w.Write("Hello")})

	router.Handle("GET /about", func(r httpfromtcp.Req, w *httpfromtcp.ResponseWriter) { w.Write("Hello from about page")})

	router.Handle("GET /service", func(r httpfromtcp.Req, w *httpfromtcp.ResponseWriter) { w.Write("service page")})
	router.ListenAndServe(":4000")
}
