package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "baidu.com")
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	err = cmd.Start() // 开始执行命令
	if err != nil {
		panic(err)
	}

	// 异步打印命令执行结果
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := stdoutPipe.Read(buf) // 从管道读取执行结果到buf中
			if err != nil {
				return
			}
			output := string(buf[:n])
			fmt.Print(output)
		}
	}()

	// 待命令执行完成，并检查命令的返回值以判断是否执行成功。
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Command finished with error: %v\n", err)
	} else {
		fmt.Println("Command finished successfully")
	}
}
