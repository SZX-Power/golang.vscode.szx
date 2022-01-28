package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//handler处理者  addr地址

func sayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(b))
	// fmt.Fprintln(w, "hello 你好")
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http serve failed,err:%v\n", err)
		return
	}
}
