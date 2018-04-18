package main

import (
	"fmt"
)

type Prototype4 struct {
	num int
	add func(p Prototype4, x int) int // private function for instance unique logic
}

func (p Prototype4) Add(x int) int { //  public function for method receiver bind and instance common logic (aka type logic)
	if x < 0 {
		return p.add(p, x)
	}

	return p.add(p, x)
}

func main() {
	p1 := Prototype4{
		num: 4,
		add: func(p Prototype4, x int) int {
			return p.num + x
		},
	}
	fmt.Println(p1.Add(4))

	p2 := Prototype4{
		num: 4,
		add: func(p Prototype4, x int) int {
			return p.num + x*x*x
		},
	}
	fmt.Println(p2.Add(4))
}

// And now you need to ask: How is this even useful?
// Why do I even want prototypes in my application?

// Cobra, the beloved command line framework uses it a lot and demonstrated a good use case.
// https://github.com/spf13/cobra#create-rootcmd
// I don't know how Cobra implements it, but I guess something similar to what I have demonstrated to you so far.
