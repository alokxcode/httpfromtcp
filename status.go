package httpfromtcp

const StatusOk int = 200
const StatusCreated int = 201
const StatusBadRequest int = 400
const StatusNotFound int = 404
const StatusInternalServerError int = 500
var ReasonPhrase string

func SetReasonPhrase(status_code int) string {
	switch status_code {
	case 200:
		ReasonPhrase = "OK"
	case 201:
		ReasonPhrase = "Created"
	case 400:
		ReasonPhrase = "Bad Request"
	case 404:
		ReasonPhrase = "Not Found"
	case 500:
		ReasonPhrase = "Internal Server Error"
	default:
		ReasonPhrase = "IDK"
	}
	return ReasonPhrase
}
