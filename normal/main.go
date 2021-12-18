package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
用原生的 net/http来实现web服务
*/

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", count)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}

}

func count(w http.ResponseWriter, r *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
}
