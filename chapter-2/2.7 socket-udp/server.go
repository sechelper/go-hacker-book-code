package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	host := "localhost"
	port := "8053"
	server, err := net.ListenPacket("udp", host+":"+port) // 监听UDP服务端口
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	buf := make([]byte, 1024)
	_, addr, err := server.ReadFrom(buf) // 接收客户端数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端[%s] <<< 服务端[%s] : %s \n", addr, server.LocalAddr(), string(buf))

	data := "欢迎来到Secself"
	_, err = server.WriteTo([]byte(data), addr) // 向客户端发送数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端[%s] <<< 服务端[%s] : %s \n", addr, server.LocalAddr(), data)

}
