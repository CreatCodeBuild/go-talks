package main

import (
	"fmt"
	"golang.org/x/tools/go/gcimporter15/testdata"
)

// Now we are going to accelerate

type Prototype3 struct {
	Add func(p Prototype3, x int) int
	num int
}

func main() {
	p := Prototype3{
		Add: func(p Prototype3, x int) int {
			return p.num + x
		},
		num: 3,
	}
	sum := p.Add(p,3)
	fmt.Println(sum)

	// Behold, do you see the problem?
	// The first p argument could be any Prototype3,
	// it does not have to be the instance of which the function is called on.

	// You need a way to hard bind the method receiver to the instance.

	// Classic JavaScript
	// Maybe, it's not the problem of JavaScript, but an inherited nature of prototypes.

	// A not too bad but really ugly solution
	// You can use the reflection lib
	// to implement something like
	/*

	func bind(instance *Prototype, bind_name string, f AnyFunctionType) a_bind_function_of_f {
		instance.bind_name = func(parameters_of_f f[1:]) {
			return f(instance, parameters_of_f)
		}
	}

	p := Prototype{
		num: 10
		f1: nil
		f2: nil
	}

	bind(p, "f1", func1)
	bind(p, "f2", func2)

	 */
	 // But using reflection produce nearly unreadable code, it is possible, but I don't even bother to try.
	 // Reflection also produce dramatically slower code due to a lots of pointer resolutions and allocations.
	 // It also gives up compile time checking of types.

	 // It does not feel idiomatic in go.
	 // You can just use JS or Python for this reason.
}

// Also, an additional problem is: what if you want to perform some common tasks on all instances of a type,
// and in the same time, keep the prototype nature,
// and also bind the method receiver?

// Can we have all of 3:

// Instance                                                                                     Code Share
// Unique                                                                                            among
// Implementation                         bind method receiver scope                             Instances
// Prototype <-------------------------------------------------------------------------------------> Class
