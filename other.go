package main

import (
	"image"
	"image/color"
	"math"
)

func drawLine(x1, y1, x2, y2 int, c color.Color) {
	dx := x2 - x1
	dy := y2 - y1
	for i := float64(0); ftoi(i) != dx; {
		canvas.Set(
			x1+ftoi(i),
			y1+ftoi(float64(dy)/float64(dx)*i),
			c,
		)

		if dx < 0 {
			i -= 0.1
			continue
		}
		i += 0.1
	}
}

func drawCircle(originX, originY int, radius float64, c color.Color, fill bool) {
	for dx := -radius; dx <= radius; dx++ {
		for dy := -radius; dy <= radius; dy++ {
			hyp := math.Sqrt(dx*dx + dy*dy)
			if math.Round(hyp) == radius || (fill && hyp < radius) {
				canvas.Set(originX+int(dx), originY+int(dy), c)
			}
		}
	}
}

func drawRect(x1, y1, x2 ,y2 int,c color.Color){
	dx := int(math.Abs(float64(x2 - x1)))
	dy := int(math.Abs(float64(y2 - y1)))
	for y := 0;y < dy; y++ {
		for x := 0; x < dx; x++ {
			canvas.Set(x1 + x, y1 + y, c);
		}
	}
}

func resetCanvas() {
	canvas = image.NewRGBA(image.Rect(0, 0, width, height))
}

func intabs(n int) int {
	return int(math.Abs(float64(n)))
}
