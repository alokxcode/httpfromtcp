package main
//
// import (
// 	"fmt"
// 	"net"
//
// 	"github.com/alokxcode/http_from_tcp/internal/hello_http"
// )
//
// func main() {
// /* 	server := hello_http.NewServer() */
//
// 	fmt.Println("Hello world")
// 	listen_addr := ":4000"
// 	listener, err := net.Listen("tcp", listen_addr)
// 	if err != nil {
// 		return
// 	}
// 	for {
// 		conn, err := listener.Accept()
// 		fmt.Println("accepted : ", conn)
// 		if err != nil {
// 			return
// 		}
// 		go hello_http.HandleConn(conn)
// 	}
// }
