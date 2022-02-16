package dmf

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
)

type Core struct {
	Request        *Request
	ResponseWriter http.ResponseWriter
	Repository     *Repository
}

func (core *Core) GetRoutes() []Route {
	return core.Repository.GetRoutes()
}

func (core *Core) HandleRequest() {
	var hasRoute = false
	var handler func(core *Core) *Response
	method := core.Request.Method
	path := core.Request.Path

	var log = &Log{Writer: os.Stdout}
	log.Info(fmt.Sprintf("%s %s", method, path))

	for _, route := range core.GetRoutes() {
		if matched, _ := regexp.MatchString(route.Pattern, path); matched == true {
			if route.HasMethod(method) {
				handler = route.Handler
				hasRoute = true
			}
		}
	}

	var response *Response
	//defer handleError(response)

	if hasRoute && handler != nil {
		response = handler(core)
	} else {
		response = &Response{}

		if !hasRoute {
			response.StatusCode = 404
			response.Content = "404 not found"
		} else if handler == nil {
			log.Error(fmt.Sprintf("failed to determine handler: path=%s, method=%s", path, method))
		}
	}

	if response != nil {
		response.Flush(core.ResponseWriter)
	} else {
		log.Error("response is nil")
	}
}

func handleError(response *Response) {
	if response == nil {
		panic("error")
	}

	if recoveryMessage := recover(); recoveryMessage != nil {
		response.StatusCode = 500
		response.Content = fmt.Sprintf("Unhandled error: %s", recoveryMessage)
	}
}

func (core *Core) String(s string) *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Content:    s,
	}
}
