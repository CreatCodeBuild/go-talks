package main

import (
	"fmt"
)

func main()  {
	for i := 0; i <= 50000; i++ {
		fmt.Println(103*103*133*i*i*i)
	}
}