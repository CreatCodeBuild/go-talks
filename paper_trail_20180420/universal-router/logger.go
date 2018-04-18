package main

import (
	"fmt"
	"time"
)

type Log struct {
	Level string
	Msg   string
}

type IntNode struct {
	Pipe chan Routable
}

func (n IntNode) Send(x Routable) {
	n.Pipe <- x
}

func (n IntNode) Receive() Routable {
	return <-n.Pipe
}

func main() {
	source := IntNode{make(chan Routable)}
	debug := IntNode{make(chan Routable)}
	info := IntNode{make(chan Routable)}
	r := Router{
		From: []PipeIn{source},
		To:   []PipeOut{debug, info},
	}

	go r.Start()

	go func() {
		for {
			source.Send(Log{
				Level: "debug",
				Msg:   "xxx",
			})

			source.Send(Log{
				Level: "info",
				Msg:   "xxx",
			})
		}
	}()

	for {
		fmt.Println(debug.Receive(), info.Receive())
		fmt.Println(debug.Receive(), info.Receive())
		time.Sleep(1 * time.Second)
	}

}

/*
	This is a infinite logger, 1 second period
 */
