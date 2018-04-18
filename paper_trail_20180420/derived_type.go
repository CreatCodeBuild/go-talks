package main

import "fmt"

type BaseType struct {
	x, y, z int
}

func (b BaseType) Sum() int {
	return b.x + b.y + b.z
}

type DerivedType struct {
	BaseType
	p1 int
	p2 int
}

func f1(v BaseType) BaseType { return v }

func f2(v Interface) Interface { return v }

type Interface interface {
	Sum() int
}

func main() {
	d := DerivedType{}
	fmt.Println(d.Sum())

	//f1(DerivedType{}) // This doesn't work

	r := f2(DerivedType{}) // This works
	fmt.Println(r)
	// DerivedType reuses the memory layout of BaseType,
	// but is  a sub scope of BaseType.

	// You can sort of implement an "inheritance" by letting BaseType implement interfaces.
	// Go proverb: Pass interface, return struct.
}
