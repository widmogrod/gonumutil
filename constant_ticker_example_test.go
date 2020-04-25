package gonumutil

import "gonum.org/v1/plot"

func ExampleNewConstantNumTicker() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.X.Tick.Marker = NewConstantNumTicker(1)
	p.Y.Tick.Marker = NewConstantNumTicker(5)
}
