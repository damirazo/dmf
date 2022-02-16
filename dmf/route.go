package dmf

type HttpMethodType string

const (
	HttpMethodGet     HttpMethodType = "GET"
	HttpMethodPost    HttpMethodType = "POST"
	HttpMethodPut     HttpMethodType = "PUT"
	HttpMethodPatch   HttpMethodType = "PATCH"
	HttpMethodDelete  HttpMethodType = "DELETE"
	HttpMethodHead    HttpMethodType = "HEAD"
	HttpMethodConnect HttpMethodType = "CONNECT"
	HttpMethodTrace   HttpMethodType = "TRACE"
	HttpMethodOptions HttpMethodType = "OPTIONS"
)

var AllHttpMethods = []HttpMethodType{
	HttpMethodGet,
	HttpMethodPost,
	HttpMethodPut,
	HttpMethodPatch,
	HttpMethodDelete,
	HttpMethodHead,
	HttpMethodConnect,
	HttpMethodTrace,
	HttpMethodOptions,
}

type Route struct {
	Pattern string
	Handler func(core *Core) *Response
	Methods []HttpMethodType
}

func (route *Route) HasMethod(method HttpMethodType) bool {
	for _, m := range route.Methods {
		if method == m {
			return true
		}
	}
	return false
}
