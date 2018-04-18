package main

type PipeOut interface {
	Send(x Routable)
}

type PipeIn interface {
	Receive() Routable
}

type Routable interface{}

/////////////////////////////////////////////////////////////////////
type Router struct {
	From []PipeIn
	To   []PipeOut
	//accept    func(data Routable) bool
	reduce    func(data []Routable) Routable
	transform func(data Routable) Routable   // A generic map X -> Y function
	apply     func(data Routable)            // If you want to apply any side effects, use with caution!
	unfold    func(data Routable) []Routable // the reverse of reduce
}

func (r Router) Start() {

	for {

		dataList := make([]Routable, len(r.From))

		//fmt.Println(dataList)

		for i, dataChan := range r.From {
			dataList[i] = dataChan.Receive()
		}

		//fmt.Println(dataList)

		var data Routable
		if r.reduce != nil {
			data = r.reduce(dataList)
		} else {
			data = dataList[0]
		}

		if r.transform != nil {
			data = r.transform(data)
		} else {
		}

		if r.apply != nil {
			r.apply(data)
		}

		var dataOut []Routable
		if r.unfold != nil {
			dataOut = r.unfold(data)
		} else {
			dataOut = []Routable{data, data}
		}

		//fmt.Println(dataOut)

		for i, to := range r.To {
			to.Send(dataOut[i])
		}
	}
}
