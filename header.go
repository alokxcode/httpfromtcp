package httpfromtcp


type Header map[string]any


func (h Header) Set(key string, value string) {
	h[key] = value	
}

