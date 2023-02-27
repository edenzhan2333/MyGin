package ginn

import (
	"fmt"
	"net/http"
	"strings"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 实现http.Handler interface{}，执行Handlerfunc
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := strings.Builder{}
	key.WriteString(req.Method)
	key.WriteString("-")
	key.WriteString(req.URL.Path)
	if handler, ok := engine.router[key.String()]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// 为engine添加路由映射，同一路由可支持不同handler
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := strings.Builder{}
	key.WriteString(method)
	key.WriteString("-")
	key.WriteString(pattern)
	engine.router[key.String()] = handler
}

// GET请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//POST请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//启动e服务，传入的engine实际作为http.Handler接口执行
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
