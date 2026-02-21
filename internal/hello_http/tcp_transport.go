package hello_http

import (
	"fmt"
	"net"
)

func HandleConn(conn net.Conn) error {
	for {
		buff := make([]byte, 20)
		n, err := conn.Read(buff)
		if err != nil {
			return err
		}
		req := string(buff[:n])
		parsed_req := Parse_Request(req)
		Response(conn, parsed_req)

		fmt.Println(parsed_req.Body)
		fmt.Println(parsed_req.Header["Content-Length"])
	}
}
