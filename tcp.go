package httpfromtcp

import (
	"fmt"
	"net"
)

func Listener(server *Server, listen_addr string) Error {

	// net.Listen returns a Listener that listens for tcp requests on a port
	listener, err := net.Listen("tcp", listen_addr)
	fmt.Println("listener is ready")
	if err != nil {
		return Error{
			Msg: "Error while listening tcp",
			Err: err,
		}
	}

	// listener should run in infinite loop to listen all incoming request - all time
	for {
		// when req come, listener will accept it
		fmt.Println("Listening on port")
		conn, err := listener.Accept()
		fmt.Println("conn accepted", conn)
		if err != nil {
			fmt.Println(Error{
				Msg: "Error while accepting request",
				Err: err,
			})
		}

		go HandleConnection(server,conn)

	}

}

func HandleConnection(server *Server, conn net.Conn) Req {
	for {
		fmt.Println("handling connection")
		// pass/set conn to server struct 
		server.res_writer.conn = conn
		// initialisation of header map
		server.res_writer.res.Header = make(map[string]any)
		// create a buffer to store the req msg temporarily to read
		buff := make([]byte, 2000)

		// read the req and write it in buffer
		n, err := conn.Read(buff)
		if err != nil {
			// 	return types.Error {
			// 		Msg : "Error while reading connection",
			// 		Err : err,
			// 	}
			return Req{}
		}
		req := string(buff[:n])
		parsed_req := Parse_Request(req)
		fmt.Println(parsed_req)
		fn := MatchRoutes(parsed_req, server)
		
		if fn == nil {
			continue
		}
		fn(*parsed_req,&server.res_writer)
		fmt.Println(server.res_writer.res.Body)

	}

}
//
// func WriteBytes(conn net.Conn, msg string) {
//
// }
