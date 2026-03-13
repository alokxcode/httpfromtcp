package httpfromtcp

import "fmt"

type Server struct {
	Listen_addr string
	Routes      map[string]HandleFunc
}

type HandleFunc func(*Req, *ResponseWriter)

// NewServer
func NewServer() Server {
	server := &Server{
		Routes: make(map[string]HandleFunc),
	}
	return *server

}

// Handle
func (server *Server) Handle(method_path string, handlefunc HandleFunc) {
	fmt.Println("routes starting to registering")
	server.Routes[method_path] = handlefunc
	fmt.Println("routes registered")
}

// ListenAndServe
func (server *Server) ListenAndServe(listen_addr string) {
	fmt.Println("Server starting")
	server.Listen_addr = listen_addr
	err := Listener(server,server.Listen_addr)
	if err.Err != nil {
		fmt.Println(err)
		return
	}
}
