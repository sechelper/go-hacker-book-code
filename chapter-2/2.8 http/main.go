package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://go-hacker-code.lab.secself.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Printf("状态：%d\n", resp.StatusCode)
	fmt.Printf("长度：%d\n", resp.ContentLength)
	fmt.Printf("内容：\n%s", string(body))
}
