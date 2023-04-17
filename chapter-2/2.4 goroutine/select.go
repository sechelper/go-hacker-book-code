package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	// 开启一个goroutine，向channel1发送消息
	go func() {
		time.Sleep(time.Second * 2)
		channel1 <- "Hello from channel 1"
	}()

	// 开启一个goroutine，向channel2发送消息
	go func() {
		time.Sleep(time.Second * 3)
		channel2 <- "Hello from channel 2"
	}()

	for {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)
		default:
			time.Sleep(time.Second)
			fmt.Println("default")
		}
	}
}
