package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"
)

// RNG Seeder
var _ = func() struct{} { rand.Seed(time.Now().UnixNano()); return struct{}{} }()

// Dimensions of screen
const width int = 800
const height int = 600
const tickWait time.Duration = 1 * time.Millisecond

var canvas *image.RGBA = image.NewRGBA(
	image.Rect(0, 0, width, height),
)

func main() {
	fmt.Println("Starting GoCanvas...")
	// Setup
	/*fls := make([]*Walker, 10, 10)
	for i, _ := range fls {
		fls[i] = NewWalker(rand.Intn(width), rand.Intn(height))
	}*/
	var seq = [][2]float64{
		{0, 2},
		{2, 0},
		{0, 0},
	}
	c := color.RGBA{rUint8(), rUint8(), rUint8(), 255}
	white := color.RGBA{255, 255, 255, 255}

	base := 0.002
	mults := make([]float64, len(seq)*2)
	go func() {
		for {
			seq = [][2]float64{
				{0, 2},
				{2, 0},
				{0, 0},
			}
			for i := range mults {
				mults[i] = []float64{-2, -1, 0, 1, 2}[rand.Intn(5)]
			}
			c = color.RGBA{rUint8(), rUint8(), rUint8(), 255}
			time.Sleep(2 * time.Second)
		}
	}()
	// /Setup

	go serve()
	fmt.Println("Server is up! CTRL-C to exit")
<<<<<<< HEAD

	// Loop
=======
	// Game Loop
>>>>>>> 0e6fb1de80b04bd0f9682b5536acd44425c71d58
	for {
		/*if rand.Int()%2000 == 0 {
			col := pickcol(fls)
			for i := 0; i < 800; i++ {
				for j := 0; j < 600; j++ {
					canvas.Set(i, j, col)
				}
			}
		}
		for _, w := range fls {
			w.Tick()
		}*/
		scatterPlot(seq, c, [2]float64{50, -50})
		time.Sleep(tickWait)
		scatterPlot(seq, white, [2]float64{50, -50})
		for i := range seq {
			seq[i][0] += base * mults[i*2]
			seq[i][1] += base * mults[i*2+1]
		}
	}
	// /Loop
}

func pickcol(a []*Walker) color.Color {
	return (*a[rand.Intn(len(a))]).c
}
