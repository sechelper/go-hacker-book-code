package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义两个通道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个 goroutine，分别从两个通道中获取数据
	go func() {
		for {
			ch1 <- ">> 1"
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			ch2 <- ">> 2"
			time.Sleep(1 * time.Second)
		}
	}()

	// 使用 select 语句非阻塞地从两个通道中获取数据
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			// 如果两个通道都没有可用的数据，则执行这里的语句
			fmt.Println(">> no message received")
			time.Sleep(1 * time.Second)
		}
	}
}
