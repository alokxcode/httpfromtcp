package types

type Req struct {
	Method  string
	Path    string
	Version string
	Header  map[string]any
	Body    string
}
