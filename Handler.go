package httpfromtcp

import "fmt"


func MatchRoutes(req *Req , server *Server) HandleFunc {
	methodPath :=  req.Method + " " + req.Path
	fmt.Println(methodPath)
	fmt.Print(server.Routes[methodPath])
	if f, ok := server.Routes[methodPath]; ok {
		fmt.Println("route matched")
		return f
	}
	return nil

}
