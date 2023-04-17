package main

import (
	"fmt"
	"go-scope/scope"
)

func main() {
	// 包名.函数
	// package.func
	scope.A()
	// package.var
	fmt.Println(scope.BStr)
	// test.b() // 无法调用
	scope.C()
	fmt.Println()
}
