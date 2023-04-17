package main

import (
	"fmt"
	"sync"
	"time"
)

var money = 100
var rw sync.RWMutex

// 余额
func balance() int {
	return money
}

// depositMoney 存钱
func depositMoney(m int) {
	rw.Lock()
	money += m
	rw.Unlock()

}

func main() {

	go depositMoney(100)
	go depositMoney(100)
	go depositMoney(100)
	go func() {
		for {
			//rw.RLock()
			time.Sleep(time.Second * 2)
			time.Sleep(time.Second)
			fmt.Print(balance(), "go1")
			//rw.RUnlock()
		}

	}()

	go func() {
		for {
			//rw.RLock()
			time.Sleep(time.Second)
			time.Sleep(time.Second)
			fmt.Print(balance(), "go2")
			//rw.RUnlock()
		}

	}()
	time.Sleep(time.Second * 1000)
}
