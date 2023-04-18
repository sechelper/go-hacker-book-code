package main

import (
	"fmt"
	scope2 "go-scope/scope"
)

var Count = 1

const User = "admin"

type Password string

var pwd Password = "admin"

func main() {
	fmt.Println(scope2.Username)
}
