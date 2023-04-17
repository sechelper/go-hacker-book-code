package main

import (
	"fmt"
	"log"
)

const secself = "secself"

func main() {
	log.Print("hello %s", secself)
	log.Fatalf("hello %s", "test") // 程序会退出
	fmt.Println(1)
}
