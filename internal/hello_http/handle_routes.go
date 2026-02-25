package hello_http

func (server *server) handle(method_path string, handlefunc handlefunc) {
	server.routes["method_path"] = handlefunc
}
