package hello_http

import (
	"strings"
)

type Req struct {
	Method  string
	Path    string
	Version string
	Header  map[string]any
	Body    string
}

func Parse_Request(req string) Req {

	fullHeader, content_body, _ := strings.Cut(req, "\r\n\r\n")
	fullHeader_slice := strings.Split(fullHeader, "\r\n")
	firstline := strings.Split(fullHeader_slice[0], " ")
	header_slice := fullHeader_slice[1:]

	header_map := make(map[string]any)

	for _, v := range header_slice {
		key, value, _ := strings.Cut(v, ": ")
		header_map[key] = value
	}

	if header_map["Content-Length"] != nil {

	}

	// fmt.Println(header_map["Host"])
	// fmt.Println(header_map["Connection"])

	parsed_req := &Req{
		Method:  firstline[0],
		Path:    firstline[1],
		Version: firstline[2],
		Header:  header_map,
		Body:    content_body,
	}
	return *parsed_req

	// fmt.Printf("secondline: %q", lines[2])
}
