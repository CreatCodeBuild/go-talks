package main

import "fmt"

// Now we are going to accelerate

type Prototype struct {
	Add func(x int) int
}

// NewProtoType uses closure to make Function Modules, if you come from JS, you're probably familiar with this term.
func NewProtoType(num int) Prototype {

	addFunc := func(x int) int {
		return num + x
	}

	return Prototype{
		Add: addFunc,
	}
}

func main() {
	p := NewProtoType(10)
	sum := p.Add(5)
	fmt.Println(sum)

	// And of course you can just use a literal
	p2 := Prototype{
		Add: func(x int) int {
			return 2 + x
		},
	}
	fmt.Println(p2.Add(5))

	// But then, the problem is,
	// How to bind the scope of the method receiver? The scope of this/self if you wish.

	// With the closure implementation, you are basically saying,
	// well, I don't want to attach any properties besides functions to my prototype,
	// and I will never want to access the method receiver (the "this" object).

	// And that's fine, as long as it solves your problem.
	// But, with this implementation,
	// you're not fully utilizing the flexibility of a prototype.
}
