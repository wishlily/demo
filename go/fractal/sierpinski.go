package fractal

import (
	"image"
	"image/color"
	"math/rand"
	"time"
)

func sierpinski_fc(img *image.RGBA, A, B, C image.Point, c []color.Color) {
	if len(c) <= 0 {
		return
	}
	triangle_1(img, A, B, C, c[0])
	D := image.Point{(A.X + B.X) / 2, (A.Y + B.Y) / 2}
	E := image.Point{(B.X + C.X) / 2, (B.Y + C.Y) / 2}
	F := image.Point{(C.X + A.X) / 2, (C.Y + A.Y) / 2}
	c1 := c[1:]
	sierpinski_fc(img, A, D, F, c1)
	sierpinski_fc(img, D, B, E, c1)
	sierpinski_fc(img, F, E, C, c1)
}

func sierpinski(img *image.RGBA, x, y int) {
	A := image.Point{x / 2, 0}
	B := image.Point{0, y}
	C := image.Point{x, y}
	c := []color.Color{
	// color.RGBA{0xEC, 0xD6, 0xC6, 255},
	// color.RGBA{0xD4, 0xDA, 0x90, 255},
	// color.RGBA{0xC1, 0xBC, 0x44, 255},
	// color.RGBA{0x63, 0x21, 0x5D, 255},
	// color.RGBA{0xB4, 0x3C, 0xAC, 255},
	// color.RGBA{0xD6, 0x85, 0xCB, 255},
	// color.RGBA{0xA1, 0x36, 0x5F, 255},
	}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	if len(c) == 0 {
		for i := 0; i < 7; i++ {
			t := r.Intn(0xFFFFFF)
			c_r := uint8(t >> 16)
			c_b := uint8(t >> 8)
			c_g := uint8(t)
			c = append(c, color.RGBA{c_r, c_b, c_g, 255})
		}
	}
	sierpinski_fc(img, A, B, C, c)
}
