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

func TestGetRoute2(t *testing.T) {
	r := newTestRouter()
	n1, ps1 := r.getRoute("GET", "/assets/file1.txt")
	ok1 := n1.pattern == "/assets/*filepath" && ps1["filepath"] == "file1.txt"
	if !ok1 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be file1.txt")
	}

	n2, ps2 := r.getRoute("GET", "/assets/css/test.css")
	ok2 := n2.pattern == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
	if !ok2 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be css/test.css")
	}

}


func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}

	if len(nodes) != 5 {
		t.Fatal("the number of routes shoule be 4")
	}
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

