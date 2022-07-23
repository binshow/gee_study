package main

import (
	"net/http"
	"testing"
)

// -------------------------------------------
// @file          : gee_test.go
// @author        : binshow
// @time          : 2022/7/23 9:11 PM
// @description   : 测试 v2 版本
// -------------------------------------------

func TestGeeV2(t *testing.T) {

	r := New()
	r.GET("/" , func(ctx *Context) {
		ctx.HTML(http.StatusOK , "<h1> hello ,binshow </h1>")
	})
	r.POST("/hello" , func(ctx *Context) {
		ctx.String(http.StatusOK , "hello %s , you are at %s\n" , "binshow" ,ctx.Path)
	})
	r.Run(":9095")

}