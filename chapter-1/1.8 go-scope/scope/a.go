package scope

import "fmt"

var BStr = "test_b"
var bstr2 = "tset_b2"

func b() {
	fmt.Println("c")
}

func A() {
	b()
}
