package dmf

import (
	"time"
)

type Repository struct {
	routes []Route
	log    *Log
}

func (r *Repository) Initialize(routes []Route, log *Log) {
	r.log = log
	r.AddRoutes(routes)

	log.Info("====================================")
	log.InfoFormat("App started at: %s", time.Now().Format("2006-01-02 15:04:05"))
	log.InfoFormat("Registered routes: %s", len(routes))
	log.Info("====================================")
}

func (r *Repository) AddRoute(pattern string, handler func(core *Core) *Response, methods ...HttpMethodType) {
	route := Route{
		Pattern: pattern,
		Handler: handler,
		Methods: methods,
	}

	r.addRoute(route)
}

func (r *Repository) AddRoutes(routes []Route) {
	for _, route := range routes {
		r.addRoute(route)
	}
	r.routes = append(r.routes, routes...)
}

func (r *Repository) addRoute(route Route) {
	if len(route.Methods) == 0 {
		route.Methods = AllHttpMethods
	}
	r.routes = append(r.routes, route)
}

func (r *Repository) GetRoutes() []Route {
	return r.routes
}

func (r *Repository) RegisterLog(log *Log) {
	r.log = log
}
