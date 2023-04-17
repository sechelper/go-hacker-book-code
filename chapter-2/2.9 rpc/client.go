package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	var result string
	err = client.Call("HelloSecself.Hello", "secself", &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
