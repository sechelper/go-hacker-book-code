package main

import (
	"fmt"
	"sync"
	"time"
)

var sw sync.WaitGroup

func worker(count int) {
	for i := 1; i < count; i++ {
		time.Sleep(500) // 暂停一秒
		fmt.Print(" ", i)
	}
	fmt.Println()
	sw.Done()
}

func main() {
	// 不使用协程
	//now1 := time.Now()
	//worker(10)
	//worker(10)
	//worker(10)
	//fmt.Println(time.Now().Sub(now1))

	// 开启三个协程同时执行
	now2 := time.Now()
	//sw.Add(3)
	go worker(10)
	sw.Add(1)
	go worker(10)
	sw.Add(1)
	go worker(10)
	sw.Add(1)
	sw.Wait()
	now3 := time.Now()
	fmt.Println(now3.Sub(now2))
}
