package parser

import (
	"fmt"
	"net"
	"github.com/alokxcode/httpfromtcp/types"
)

type ResponseWriter struct {
	res *types.Res
	conn net.Conn
}

func (rw *ResponseWriter) Write(body string) {
	rw.res.Body = body 
	rw.res.StatusCode = 200
} 

func Response(conn net.Conn, req_firstLine types.Req) {
	var msg string
	var s_code int
	var s_msg string

	switch req_firstLine.Path {
	case "/":
		msg = "Hello World!"
		s_code = 200
		s_msg = "OK"
	case "/about":
		msg = "About Page"
		s_code = 200
		s_msg = "OK"
	default:
		msg = "404 Page not found"
		s_code = 404
		s_msg = "ERROR"

	}
	msg_length := len(msg)

	res := fmt.Sprintf("HTTP/1.1 %v %v\r\nContent-Length: %v\r\n\r\n%v", s_code, s_msg, msg_length, msg)
	conn.Write([]byte(res)) 
}
