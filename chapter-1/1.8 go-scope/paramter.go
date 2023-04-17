package main

import "fmt"

var p string = "p_string"

const p_const string = "p_const"

type uuid string

func print(id uuid) {
	var p_within = "p_within"
	fmt.Println(p_within)
	fmt.Println(id)
	fmt.Println(p_const)
}

func main() {
	// ed1bd8ec-e971-44ab-8075-9264388be1d1
	print("67315c67-7d75-4b84-92bb-fbec9319185c")
}
