package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	host := "localhost"
	port := "8080"
	server, err := net.Listen("tcp", host+":"+port) // 监听机器端口
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	connection, err := server.Accept() // 等待客户端连接
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端[%s] 连接到 服务端[%s]\n", connection.RemoteAddr(), connection.LocalAddr())

	buffer := make([]byte, 1024)
	length, err := connection.Read(buffer) // 读取客户端发送的数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端[%s] >>> 服务端[%s] : %s \n", connection.RemoteAddr(), connection.LocalAddr(),
		string(buffer[:length]))

	data := "欢迎来到Secself"
	_, err = connection.Write([]byte(data)) // 向客户端发送数据
	fmt.Printf("客户端[%s] <<< 服务端[%s] : %s \n", connection.RemoteAddr(), connection.LocalAddr(),
		data)

	connection.Close()

}
