package main

import "log"

func main() {
	// 2023/04/12 13:15:52.555533 custom.go:8: hello secself
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	log.Printf("hello %s", "secself")
}
