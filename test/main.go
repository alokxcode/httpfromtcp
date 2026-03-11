package main

import "github.com/alokxcode/httpfromtcp"

func main() {
	router := httpfromtcp.NewServer()

	router.Handle("GET /", func(r httpfromtcp.Req, w *httpfromtcp.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpfromtcp.StatusBadRequest)

		w.Header().Set("Connection","keep-alive")
		w.Write("Hello")
	})

	router.Handle("GET /about", func(r httpfromtcp.Req, w *httpfromtcp.ResponseWriter) { w.Write("Hello from about page")})

	router.Handle("GET /service", func(r httpfromtcp.Req, w *httpfromtcp.ResponseWriter) { w.Write("service page")})
	router.ListenAndServe(":4000")
}
