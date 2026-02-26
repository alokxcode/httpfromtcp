package tcp

import (
	"fmt"
	"net"

	"github.com/alokxcode/httpfromtcp/types"
)

func Listener(listen_addr string) types.Error {

	// net.Listen returns a Listener that listens for tcp requests on a port
	listener, err := net.Listen("tcp", listen_addr)
	if err != nil {
		return types.Error{
			Msg: "Error while listening tcp",
			Err: err,
		}	
	}

	// listener should runs in infinite loop to listen all incoming request - all time
	for {
		// when req come, listener will accept it
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(types.Error {
				Msg: "Error while accepting request",
				Err : err,
			})
		}

		go HandleConnection(conn)

	}

}
