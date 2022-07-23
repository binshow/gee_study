package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// -------------------------------------------
// @file          : context.go
// @author        : binshow
// @time          : 2022/7/23 8:48 PM
// @description   : 封装web context
// -------------------------------------------



type H map[string]interface{}

type Context struct {
	// 请求 和 响应
	Writer http.ResponseWriter
	Req    *http.Request

	// req info
	Path   string
	Method string

	// resp info
	StatusCode int


}

func NewContext(writer http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:     writer,
		Req:        req,
		Path:       req.URL.Path,
		Method:     req.Method,
	}
}

//解析表单数据
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//设置resp的状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//设置resp的请求头的信息
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key , value)
}

//设置resp的不同格式


func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Context-Type" , "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format , values...)))
}


func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}






