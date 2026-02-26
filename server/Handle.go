package server 

type handlefunc func()
func (server *Server) Handle(method_path string, handlefunc handlefunc) {
	server.Routes["method_path"] = handlefunc
}
