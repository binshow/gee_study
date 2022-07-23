package main

import (
	"fmt"
	"net/http"
)

// -------------------------------------------
// @file          : router.go
// @author        : binshow
// @time          : 2022/7/23 8:57 PM
// @description   : 路由管理
// -------------------------------------------

type HandlerFunc func(ctx *Context)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

//添加路由

func (r *Router) addRoute(method, pattern string, handler HandlerFunc) {
	fmt.Printf("Route %4s - %s" , method , pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 处理
func (r *Router) handler(c *Context) {
	key := c.Method + "-" + c.Path
	if handler , ok := r.handlers[key]; ok{
		handler(c)
	}else {
		c.String(http.StatusNotFound , "404 not found:%s\n" , c.Path)
	}
}




