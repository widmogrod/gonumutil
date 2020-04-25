// Package provides utility functions for [GoNum](https://github.com/gonum)
package gonumutil

import (
	"fmt"
	"gonum.org/v1/plot"
)

// NewConstantNumTicker simplify change of scale of the X and Y axis to have constant distance between markers.
// step represents what distance between marks on a axis you want to preserve.
func NewConstantNumTicker(step uint) plot.TickerFunc {
	return NewConstantTicker("%1.0f", float64(step))
}

// NewConstantNumTicker simplify change of scale of the X and Y axis to have constant distance between markers.
// label give you possibility to customize label. Accepts `Sprintf` formatting specification.
// step represents what distance between marks on a axis you want to preserve.
func NewConstantTicker(label string, step float64) plot.TickerFunc {
	return plot.TickerFunc(func(min, max float64) []plot.Tick {
		var res = make([]plot.Tick, 0)

		val := min
		for val <= max {
			res = append(res, plot.Tick{
				Value: val,
				Label: fmt.Sprintf(label, val),
			})
			val += step
		}
		return res
	})
}
