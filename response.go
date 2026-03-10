package httpfromtcp

import (
	"fmt"
	"net"
)

type Res struct {
	StatusCode int
	Header     map[string]string
	Body       string
}

type ResponseWriter struct {
	res  Res
	conn net.Conn
}

func (rw *ResponseWriter) Write(body string) {
	var s_msg string


	rw.res.Body = body
	rw.res.StatusCode = 200
	s_msg = "OK"

	msg_length := len(rw.res.Body)

	res := fmt.Sprintf("HTTP/1.1 %v %v\r\nContent-Length: %v\r\n\r\n%v", rw.res.StatusCode, s_msg, msg_length, rw.res.Body)
	rw.conn.Write([]byte(res))
}
//
// func Response(conn net.Conn, req_firstLine *Req) {
// 	switch req_firstLine.Path {
// 	case "/":
// 		msg = "Hello World!"
// 		s_code = 200
// 		s_msg = "OK"
// 	case "/about":
// 		msg = "About Page"
// 		s_code = 200
// 		s_msg = "OK"
// 	default:
// 		msg = "404 Page not found"
// 		s_code = 404
// 		s_msg = "ERROR"
//
// 	}
// }
