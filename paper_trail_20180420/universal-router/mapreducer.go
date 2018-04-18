package main

import "fmt"

type IntNode struct {
	Pipe chan Routable
}

func (n IntNode) Send(x Routable) {
	n.Pipe <- x
}

func (n IntNode) Receive() Routable {
	return <-n.Pipe
}

type StringPipe struct {
	Pipe chan string
}

func (n StringPipe) Send(x Routable) {
	if str, ok := x.(string); ok {
		n.Pipe <- str
	}
}

func (n StringPipe) Receive() Routable {
	return <-n.Pipe
}


func main() {
	stringPipe := StringPipe{make(chan string)}
	chan1 := IntNode{make(chan Routable)}
	chan2 := IntNode{make(chan Routable)}

	finals := make([]IntNode, 50000)
	finalsPipeOut := make([]PipeOut, 50000)
	for i := 0; i < len(finals); i++ {
		finals[i] = IntNode{make(chan Routable)}
		finalsPipeOut[i] = finals[i]
	}

	stringMapInt := Router{
		transform: func(data Routable) Routable {
			if data, ok := data.(string); ok {
				return len(data)
			}
			return 0
		},
		From: []PipeIn{stringPipe},
		To:   []PipeOut{chan1},
	}

	mul := Router{
		reduce: func(data []Routable) Routable {
			sum := 0
			for _, d := range data {
				if num, ok := d.(int); ok {
					sum += num
				}
			}
			return sum
		},
		unfold: func(data Routable) []Routable {
			if sum, ok := data.(int); ok {
				var ret []Routable
				for i := 1; i <= 50000; i++ {
					ret = append(ret, sum*sum*sum*i*i*i)
				}
				return ret
			}
			return nil
		},
		From: []PipeIn{chan1, chan2},
		To:   finalsPipeOut,
	}

	go stringMapInt.Start()
	go mul.Start()

	stringPipe.Send("123")
	chan2.Send(100)

	for _, pipe := range finals {
		fmt.Println(pipe.Receive())
	}
}
