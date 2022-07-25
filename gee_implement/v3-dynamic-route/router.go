package main

import (
	"net/http"
	"strings"
)

// -------------------------------------------
// @file          : router.go
// @author        : binshow
// @time          : 2022/7/23 8:57 PM
// @description   : 路由管理
// -------------------------------------------

type HandlerFunc func(ctx *Context)

type router struct {
	roots    map[string]*node		// key 为不同的方法，value 为不同的方法对应的前缀树的根节点
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{
		roots: make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern , "/")
	parts := make([]string , 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts , item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}


// 添加路由
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	// 判断当前方法是否存在 root 节点
	_ , ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	// 将当前路径插入到 root 节点下
	key := method + "-" + pattern
	parts := parsePattern(pattern)
	r.roots[method].insert(pattern , parts , 0)
	r.handlers[key] = handler
}

// 查找路由
func (r *router) getRoute(method string , path string) (*node , map[string]string) {
	root , ok := r.roots[method]
	if !ok {
		return nil , nil
	}

	params := make(map[string]string)
	searchParts := parsePattern(path)
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}

			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:] , "/")
				break
			}
		}
		return n , params
	}
	return nil , nil
}

// 处理
func (r *router) handler(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	}else {
		c.String(http.StatusNotFound , "404 not found:%s\n" , c.Path)
	}
}




