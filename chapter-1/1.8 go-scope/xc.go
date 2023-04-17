package main

import "fmt"

type age int

var c age = 12

const b_const = "test_b_string"

var a = "string_var"

func test(a string) {
	const b_const = "test2"

	fmt.Println(b_const)
}

func main() {
	test("a_test_string")
}
