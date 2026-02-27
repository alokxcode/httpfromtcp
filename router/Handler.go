package router

import (
	"github.com/alokxcode/httpfromtcp/parser"
	"github.com/alokxcode/httpfromtcp/server"
	"github.com/alokxcode/httpfromtcp/types"
)
func Handler(req types.Req, server server.Server) {
	methodPath :=  req.Method + " " + req.Path
	if f, ok := server.Routes[methodPath]; ok {
		f(req, &parser.ResponseWriter )
	}
}
