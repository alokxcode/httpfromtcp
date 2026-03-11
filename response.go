package httpfromtcp

import (
	"fmt"
	"net"
	"strings"
)

	

type Res struct {
	StatusCode int
	Header    Header 
	Body       string
}

type ResponseWriter struct {
	res  Res
	conn net.Conn
}

func (rw *ResponseWriter) Write(body string) {
	var sb strings.Builder

	rw.res.Body = body
	SetReasonPhrase(rw.res.StatusCode)
	msg_length := len(rw.res.Body)
	rw.res.Header["Content-Length"] = msg_length

	for key,value := range rw.res.Header {
		sb.WriteString(fmt.Sprintf("%v: %v\r\n", key, value))
	}

	fmt.Println("content type", rw.res.Header["Content-Type"])
	res := fmt.Sprintf("HTTP/1.1 %v %v\r\n%v\r\n%v", rw.res.StatusCode, ReasonPhrase,sb.String(), rw.res.Body)
	rw.conn.Write([]byte(res))
}
func (rw *ResponseWriter) Header() Header {
	return rw.res.Header
}  

func (rw *ResponseWriter) WriteHeader(status_code int) {
	rw.res.StatusCode = status_code
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
