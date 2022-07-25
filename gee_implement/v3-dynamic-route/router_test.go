package main

import (
	"fmt"
	"testing"
)

// -------------------------------------------
// @file          : router_test.go
// @author        : binshow
// @time          : 2022/7/25 7:45 PM
// @description   : 测试用前缀树实现的 router
// -------------------------------------------

func newTestRouter() *router {
	r := NewRouter()
	r.addRoute("GET" , "/" , nil)
	r.addRoute("GET" , "/hello/:name" , nil)
	r.addRoute("GET" , "/hello/b/c" , nil)
	r.addRoute("GET" , "/hi/:name" , nil)
	r.addRoute("GET" , "/assets/*filepath" , nil)
	return r
}

func TestGetRouter(t *testing.T) {
	r := newTestRouter()
	n , ps := r.getRoute("GET" , "/hello/binshow")

	if n == nil {
		t.Fatal("nil should not be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "binshow" {
		t.Fatal("name should be equal to 'binshow'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
}


func TestParsePattern(t *testing.T) {
	pattern := "/hello/binshow/1"
	fmt.Println(parsePattern(pattern))

	pattern = "hello/a/b"
	fmt.Println(parsePattern(pattern))

	pattern = "/hello/*/ab"
	fmt.Println(parsePattern(pattern))	// [hello *]
}

func TestAddRouter(t *testing.T) {
	router := NewRouter()
	router.addRoute("GET" , "/hello/binshow" , func(ctx *Context){
		fmt.Println(ctx.Method)
	})

	roots := router.roots
	for key, node := range roots {
		fmt.Println("key = " , key)
		fmt.Println("node.pattern = " , node.pattern)
	}
	fmt.Println(router)
}

