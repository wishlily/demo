package fractal

import (
	"image"
	"image/color"
	"math"
)

func line_len(x1, y1, x2, y2 int) float64 {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	l := math.Sqrt(dx*dx + dy*dy)
	return l
}

func point(img *image.RGBA, x, y int, c color.Color) {
	img.Set(x, y, c)
}

func line(img *image.RGBA, x1, y1, x2, y2 int, c color.Color) {
	delta_x := x2 - x1
	delta_y := y2 - y1
	if delta_x == 0 && delta_y == 0 {
		img.Set(x1, y1, c)
		return
	}

	abs_delta_x := math.Abs(float64(delta_x))
	abs_delta_y := math.Abs(float64(delta_y))

	x := float64(x1)
	y := float64(y1)
	if abs_delta_y > abs_delta_x {
		dx := float64(delta_x) / abs_delta_y
		dy := float64(delta_y) / abs_delta_y
		for i := 0; i <= int(abs_delta_y); i++ {
			img.Set(int(x+0.5), int(y), c)
			x += dx
			y += dy
		}
	} else {
		dx := float64(delta_x) / abs_delta_x
		dy := float64(delta_y) / abs_delta_x
		for i := 0; i <= int(abs_delta_x); i++ {
			img.Set(int(x), int(y+0.5), c)
			x += dx
			y += dy
		}
	}
}
