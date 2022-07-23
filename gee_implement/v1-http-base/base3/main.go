package main

import (
	"fmt"
	"net/http"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/23 3:12 PM
// @description   : 使用自己实现的 Engine
// -------------------------------------------

func main() {
	r := New()
	r.GET("/path" , func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w , "url.path = %q\n" , r.URL.Path)
	})
	fmt.Println("router map = " , r.router) // router map =  map[GET-/path:0x10058b6b0]
	r.Run(":9997")
}