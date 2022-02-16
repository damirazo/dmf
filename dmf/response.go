package dmf

import (
	"io"
	"net/http"
)

type Response struct {
	StatusCode int
	Content    string
}

func (r *Response) Write(s string) {
	r.Content += s
}

func (r *Response) Flush(writer http.ResponseWriter) {
	io.WriteString(writer, r.Content)
}
