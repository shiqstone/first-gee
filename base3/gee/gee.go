package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (engine Engine) addRoute(method string, pattern string, handle HandleFunc) {
	key := method + "-" + pattern
	engine.router[key] = handle
}

func (engine *Engine) GET(pattern string, handle HandleFunc) {
	engine.addRoute("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handle HandleFunc) {
	engine.addRoute("POST", pattern, handle)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handle, ok := engine.router[key]; ok {
		handle(w, req)
	}else{
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}