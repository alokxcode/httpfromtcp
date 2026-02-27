package main

import (
	"fmt"

	"github.com/alokxcode/httpfromtcp/parser"
	"github.com/alokxcode/httpfromtcp/server"
	"github.com/alokxcode/httpfromtcp/types"
)
 
func main () {
	fmt.Println("server starting")
	router := server.NewServer()
	router.Handle("GET /", func(req types.Req, res parser.ResponseWriter) {})
	router.Listen(":4000")
}
