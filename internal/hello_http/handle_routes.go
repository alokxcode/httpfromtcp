package hello_http

func (server *Server) Handle(method_path string, handleFunc HandleFunc) {
	server.Routes["method_path"] = handleFunc
}
