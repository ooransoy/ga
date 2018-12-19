package main

import (
	"image/color"
	"math"
)

type ConnectMode uint8

const (
	None ConnectMode = iota
	Open
	Closed
)

func scatterPlot(d [][2]float64, c color.Color, scale [2]float64) {
	for i, p := range d {
		drawCircle(
			width/2+ftoi(p[0]*scale[0]),
			height/2+ftoi(p[1]*scale[1]),
			3, c, true,
		)
		drawLine(
			width/2+ftoi(p[0]*scale[0]),
			height/2+ftoi(p[1]*scale[1]),
			width/2+ftoi(d[(i+1)%len(d)][0]*scale[0]),
			height/2+ftoi(d[(i+1)%len(d)][1]*scale[1]),
			c,
		)
	}
}

func ftoi(f float64) int {
	return int(math.Floor(f))
}
