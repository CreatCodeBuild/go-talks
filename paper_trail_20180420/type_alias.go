package main

import "fmt"

type Int int


func main() {
	x := 1

	//fmt.Println(x == Int(1)) //not the same type, syntax error

	fmt.Println(Int(x) == Int(1))

	// Type alias / Type renaming defines a different type with the same memory layout.
	// They can be converted back n forth back,
	// but not inter changeable
}
