package main

import (
	"log"
	"os"
)

func main() {

	// 打开日志文件，不存在则创建一个新的
	file, err := os.OpenFile("custom.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.Lshortfile|log.Ldate|log.Lmicroseconds)
	logger.Println("hello secself")
}
