package main

import (
	"dmf/dmf"
	"fmt"
	"net/http"
	"os"
)

var repository *dmf.Repository

var routes = []dmf.Route{
	{Pattern: "/", Handler: func(core *dmf.Core) *dmf.Response {
		return core.String("hello world")
	}},
	{Pattern: "/ping", Handler: func(core *dmf.Core) *dmf.Response {
		return core.String("pong")
	}},
}

func main() {
	var log = &dmf.Log{Writer: os.Stdin}
	repository = &dmf.Repository{}
	repository.Initialize(routes, log)

	http.HandleFunc("/", handler)
	var err = http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("Server start error: %s", err))
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	core := &dmf.Core{}
	core.Request = dmf.NewRequest(request)
	core.ResponseWriter = writer
	core.Repository = repository
	core.HandleRequest()
}
