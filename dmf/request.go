package dmf

import "net/http"

type Request struct {
	R         *http.Request
	Method    HttpMethodType
	Path      string
	UrlParams map[string]string
}

func NewRequest(r *http.Request) *Request {
	request := &Request{
		R:      r,
		Method: HttpMethodType(r.Method),
		Path:   r.URL.Path,
	}

	return request
}
