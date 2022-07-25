package main

import "net/http"

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/25 9:19 PM
// @description   : 测试动态路由的使用
// -------------------------------------------

func main()  {
	r := New()
	r.GET("/", func(c *Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *Context) {
		c.JSON(http.StatusOK, H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}