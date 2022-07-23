package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/23 2:46 PM
// @description   :	go 标准库实现 web 服务器
// -------------------------------------------

func main() {
	http.HandleFunc("/" , indexHandler)
	http.HandleFunc("/hello" , helloHandler)
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/body/multi", getBodyIsNil)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/form", form)

	// handler is nil , 第二个参数是一个 Handler interface 类型
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}


// get the request header return to resp
func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w,"Header[%q] = %q \n" , k , v)
	}
}


// echoes req url path
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w , "url.path = %q\n" , r.URL.Path)
}


func readBodyOnce(w http.ResponseWriter, r *http.Request)  {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 记住要返回，不然就还会执行后面的代码
		return
	}
	// 类型转换，将 []byte 转换为 string
	fmt.Fprintf(w, "read the data: %s\n", string(body))

	// 尝试再次读取，啥也读不到，但是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		fmt.Fprintf(w, "read the data one more time got error: %v", err)
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s] and read data length %d\n", string(body), len(body))
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "query is %v\n", values)
}

func wholeUrl(w http.ResponseWriter, r *http.Request)  {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}


func form(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		fmt.Fprint(w, "GetBody is nil\n")
	} else {
		fmt.Fprintf(w, "GetBody not nil\n")
	}
}