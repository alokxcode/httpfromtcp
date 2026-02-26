package server

import (
	"fmt"

	"github.com/alokxcode/httpfromtcp/tcp"
)

func (server *Server) Listen(listen_addr string) {
	server.Listen_addr = listen_addr
	err := tcp.Listener(server.Listen_addr)
	if err.Err != nil{
		fmt.Println(err)
		return
	}
}
