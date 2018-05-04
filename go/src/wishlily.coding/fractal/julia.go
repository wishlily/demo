package fractal

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

/*
 * fc = z^2 + c
 * 给定一个 c 值，所有使 fc 迭代 n 次不发散的 z 值集合
 */
const (
	_JULIA_R = 2.0
)

func julia_fc(zx, zy, cx, cy float64) (float64, float64) {
	z_x := zx*zx - zy*zy + cx
	z_y := 2*zx*zy + cy
	return z_x, z_y
}

/*
 * 在循环内发散速度，灰度表示（255）
 * 越早发散图像越浅
 */
func julia_divergent_grey(zx, zy, cx, cy float64) int {
	for i := 255; i >= 0; i -= 3 {
		zx, zy = julia_fc(zx, zy, cx, cy)
		if math.Sqrt(zx*zx+zy*zy) > _JULIA_R {
			return i
		}
	}
	return 0
}

func julia(img *image.RGBA, limit int) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	// c = [-1, 1]
	cx := r.Float64()*2 - 1
	cy := r.Float64()*2 - 1
	fmt.Println(cx, cy)

	for i := 0; i < limit; i++ {
		zx := float64(4*i)/float64(limit) - _JULIA_R
		for j := 0; j < limit; j++ {
			zy := float64(4*j)/float64(limit) - _JULIA_R
			gray := uint8(julia_divergent_grey(zx, zy, cx, cy))
			point(img, i, j, color.Gray{gray})
		}
	}
}

func mandelbrot(img *image.RGBA, limit int) {
	// z = 0
	zx := 0.0
	zy := 0.0

	for i := 0; i < limit; i++ {
		cx := float64(3*i)/float64(limit) - 2
		for j := 0; j < limit; j++ {
			cy := float64(3*j)/float64(limit) - 1.5
			gray := uint8(julia_divergent_grey(zx, zy, cx, cy))
			point(img, i, j, color.Gray{gray})
		}
	}
}
