package httpfromtcp 

type Server struct {
	Listen_addr string
	Routes      map[string]func()
}

type HandleFunc func()

func NewServer() Server {
	server := &Server{}
	return *server

}



