package main

import (
	"fmt"
	"os"
)

func OpenFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		//log.Fatal(err)
		panic(err) // try catch
		//return err
	}
	fmt.Println("c")
	// 开始操作文件
	fmt.Println(file)
	//...
	return nil
}

func main() {
	defer func() {
		r := recover() // 捕获异常
		if r != nil {
			fmt.Println("recover() test,", r) // 处理文件读取异常错误
		}
	}()

	err := OpenFile("secself.log")
	if err != nil {
		fmt.Println(err)
		return
	}
}
