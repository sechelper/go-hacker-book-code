package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	host := "localhost"
	port := "8053"
	server, err := net.ResolveUDPAddr("udp", host+":"+port) // 构造服务端的地址

	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, server) // 类似于UDP网络拨号
	if err != nil {
		log.Fatal(err)
	}

	//close the connection
	defer conn.Close()

	data := "我是客户端"
	_, err = conn.Write([]byte(data)) // 向服务端发送数据

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端[%s] >>> 服务端[%s] : %s\n", conn.LocalAddr(), conn.RemoteAddr(), data)

	received := make([]byte, 1024)
	_, err = conn.Read(received) // 接收服务端数据

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("客户端[%s] <<< 服务端[%s] : %s\n", conn.LocalAddr(), conn.RemoteAddr(), received)

}
