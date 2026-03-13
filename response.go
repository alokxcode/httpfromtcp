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

	// defines string builder
	var sb strings.Builder

	// sets body in response struct
	rw.res.Body = body

	// sets reason phrase based on status code 
	SetReasonPhrase(rw.res.StatusCode)

	// sets body length to res header 
	msg_length := len(rw.res.Body)
	rw.res.Header["Content-Length"] = msg_length
	
	// gives a http formated header string
	for key,value := range rw.res.Header {
		sb.WriteString(fmt.Sprintf("%v: %v\r\n", key, value))
	}

	// gives a http formated response string
	res := fmt.Sprintf("HTTP/1.1 %v %v\r\n%v\r\n%v", rw.res.StatusCode, ReasonPhrase,sb.String(), rw.res.Body)

	// writes the response to the connection
	rw.conn.Write([]byte(res))
}
func (rw *ResponseWriter) Header() Header {
	return rw.res.Header
}  

func (rw *ResponseWriter) WriteHeader(status_code int) {
	rw.res.StatusCode = status_code
}


