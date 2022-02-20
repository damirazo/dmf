package dmf

import (
	"fmt"
	"regexp"
	"strings"
)

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
	Pattern         string
	CompiledPattern *regexp.Regexp
	Methods         []HttpMethodType
	Handler         func(core *Core) *Response
}

func (route *Route) GetPattern() *regexp.Regexp {
	if route.CompiledPattern != nil {
		return route.CompiledPattern
	}

	var pattern = route.Pattern
	var segment = "[a-z0-9_\\-]+?"
	var reArgs = fmt.Sprintf(`<(%s)>`, segment)
	re, _ := regexp.Compile(reArgs)
	matched := re.FindAllString(pattern, -1)
	if matched != nil {
		for _, match := range matched {
			var namedPattern = fmt.Sprintf(`(?P%s%s)`, match, segment)
			pattern = strings.Replace(pattern, match, namedPattern, 1)
		}
	}
	route.CompiledPattern, _ = regexp.Compile(fmt.Sprintf("^%s$", pattern))

	return route.CompiledPattern
}

func (route *Route) HasMethod(method HttpMethodType) bool {
	for _, m := range route.Methods {
		if method == m {
			return true
		}
	}
	return false
}

func (route *Route) Match(path string, method HttpMethodType) (bool, map[string]string) {
	var pattern = route.GetPattern()
	context := make(map[string]string)
	var isMatched = pattern.MatchString(path)
	if isMatched && route.HasMethod(method) {
		match := pattern.FindStringSubmatch(path)
		for i, name := range pattern.SubexpNames() {
			if i != 0 && name != "" {
				context[name] = match[i]
			}
		}

		return true, context
	}

	return false, nil
}
