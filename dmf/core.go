package dmf

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
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
	var urlParams map[string]string
	method := core.Request.Method
	path := CleanPath(core.Request.Path)

	var log = &Log{Writer: os.Stdout}
	log.Info(fmt.Sprintf("%s %s", method, path))

	for _, route := range core.GetRoutes() {
		isMatched, ctx := route.Match(path, method)
		if isMatched {
			handler = route.Handler
			hasRoute = true
			urlParams = ctx
			break
		}
	}

	var response *Response
	core.Request.UrlParams = urlParams
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

func (core *Core) String(s string) *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Content:    s,
	}
}

// ApplicationRoot Путь до директории приложения
func (core *Core) ApplicationRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

// Template рендеринг указанного шаблона
func (core *Core) Template(templateName string, context map[string]interface{}) (string, error) {
	var err error

	t := template.New(templateName)

	d, err := core.ApplicationRoot()
	if err != nil {
		return "", nil
	}

	path := filepath.Join(d, "templates", templateName)
	t, err = t.ParseFiles(path)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, context); err != nil {
		return "", err
	}

	query := tpl.String()

	return query, nil

}

// Утилиты

func handleError(response *Response) {
	if response == nil {
		panic("error")
	}

	if recoveryMessage := recover(); recoveryMessage != nil {
		response.StatusCode = 500
		response.Content = fmt.Sprintf("Unhandled error: %s", recoveryMessage)
	}
}
