package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	go func() {
		time.Sleep(3 * time.Second) // 等待三秒
		channel <- "Hello secself!" // 向通道插入消息
	}()

	fmt.Println("开始等待消息...")
	message := <-channel
	fmt.Println(message)
}
