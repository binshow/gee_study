package main

import (
	"fmt"
	"net/http"
)

// -------------------------------------------
// @file          : gee.go
// @author        : binshow
// @time          : 2022/7/23 9:02 PM
// @description   : 框架入口
// -------------------------------------------

type Engine struct {
	router  *Router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter,  r *http.Request)  {
	fmt.Println("engine serverHTTP start")
	c := NewContext(w , r)
	engine.router.handler(c)
}


func (engine *Engine) addRoute(method string , pattern string , handler HandlerFunc) {
	engine.router.addRoute(method, pattern , handler)
}

//===  包装一些常用的方法

func (engine *Engine) GET(pattern string , handler HandlerFunc) {
	engine.addRoute("GET" , pattern , handler)
}


func (engine *Engine) POST(pattern string , handler HandlerFunc) {
	engine.addRoute("POST" , pattern , handler)
}


func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr , engine)
}



