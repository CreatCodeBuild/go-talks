package main

import "fmt"

// Now we are going to accelerate

type Prototype2 struct {
	Add func(x int) int
	num int
}

func InitPrototype2(p Prototype2) Prototype2 {

	p.Add = func(x int) int {
		return p.num + x
	}

	return p
}

func main() {
	p := Prototype2{num: 11}
	p = InitPrototype2(p)
	sum := p.Add(5)
	fmt.Println(sum)

	// Notice, that both this and the previous implementation provide private "fields"
	// However, this implementation provides the option to make "num" public -> "Num"

	// It's also yours to decide whether to make these closure function receive pointers,
	// depending on whether you want to construct a new prototype or modify the existing one.

	// This implementaation also makes your prototype Type bigger.

	// But, you might feel this use of closure is totally unnecessary,
	// or, your machine does not support heap (embeded system with no heap)?

	// Note:
	// A closure cannot be (easily) implemented purely on stack,
	// unless you pre allocate a chunk of stack at the top to use as a heap,
	// But that makes the heap not dynamic.

	// Let's take a look at another implementation
}
