package server

import (
	"github.com/alokxcode/httpfromtcp/parser"
	"github.com/alokxcode/httpfromtcp/types"
)

type Server struct {
	Listen_addr string
	Routes      map[string]func()
}

type HandleFunc func(types.Req, parser.ResponseWriter)

func NewServer() Server {
	server := &Server{}
	return *server

}



