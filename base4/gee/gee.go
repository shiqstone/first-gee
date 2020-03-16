package gee

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandleFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handle HandleFunc) {
	engine.router.addRoute(method, pattern, handle)
}

func (engine *Engine) GET(pattern string, handle HandleFunc) {
	engine.addRoute("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handle HandleFunc) {
	engine.addRoute("POST", pattern, handle)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}