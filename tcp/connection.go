 package tcp

import (
	"fmt"
	"net"
/* 	"github.com/alokxcode/httpfromtcp/types" */
	"github.com/alokxcode/httpfromtcp/parser"
)

func HandleConnection(conn net.Conn) {
	for {
		// create a buffer to store the req msg temporarily to read
		buff := make([]byte, 2000)

		// read the req and write it in buffer
		n, err := conn.Read(buff)
		if err != nil {
		// 	return types.Error {
		// 		Msg : "Error while reading connection",
		// 		Err : err,
		// 	}
		return 
		}
		req := string(buff[:n])
		fmt.Println(req)
		parsed_req := parser.Parse_Request(req)
		parser.Response(conn, parsed_req)

		fmt.Println(parsed_req.Body)
		fmt.Println(parsed_req.Header["Content-Length"])
	}

}
