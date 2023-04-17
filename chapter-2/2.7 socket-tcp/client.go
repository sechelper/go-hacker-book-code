package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	host := "localhost"
	port := "8080"

	connection, err := net.Dial("tcp", host+":"+port) // 连接到服务端
	if err != nil {
		panic(err)
	}

	data := "我是客户端"
	_, err = connection.Write([]byte(data)) // 向服务端发送数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端[%s] >>> 服务端[%s] : %s\n", connection.LocalAddr(), connection.RemoteAddr(), data)

	buffer := make([]byte, 1024)
	length, err := connection.Read(buffer) // 读取服务端发送的数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端[%s] <<< 服务端[%s] : %s\n", connection.LocalAddr(), connection.RemoteAddr(), string(buffer[:length]))

	connection.Close() // 关闭与客户端的连接
}
