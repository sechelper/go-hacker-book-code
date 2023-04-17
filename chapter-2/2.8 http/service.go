package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Secself!")
}

func main() {
	// 注册http处理函数
	http.HandleFunc("/", handler)

	// 启动http服务端
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
