package main

import (
	"fmt"

	"github.com/alokxcode/httpfromtcp/server"
)
 
func main () {
	fmt.Println("server starting")
	router := server.NewServer()
	router.Listen(":4000")
}
