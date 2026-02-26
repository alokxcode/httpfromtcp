package server

import (
	"fmt"

	"github.com/alokxcode/httpfromtcp/tcp"
)

func (server *Server) Listen() {
	err := tcp.Listener(server.Listen_addr)
	if err.Err != nil{
		fmt.Println(err)
		return
	}
}
