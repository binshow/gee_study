package main


import (
	"fmt"
	"net/http"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/23 2:58 PM
// @description   : 自己实现 Handler interface 作为框架入口
// -------------------------------------------

type Engine struct {}

// 实现了 Handler 接口， 拦截了所有的 http请求，有了统一的控制入口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// 所有的 http 请求都会进行自己实现的 engine
	fmt.Fprintf(w , "you own engine start....\n")

	// 得在这里进行注册路由了
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w , "this path is %q\n" , r.URL.Path)
	case "/hello":
		fmt.Fprintf(w, "hello ,you are handsome")
	default:
		fmt.Fprintf(w , "404 not found")
	}

}


func main() {
	http.HandleFunc("/" , helloHandler)
	engine := &Engine{}
	err := http.ListenAndServe(":9998", engine)	//http://localhost:9998/
	if err != nil {
		fmt.Println(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w , "hello, your own handler")
}