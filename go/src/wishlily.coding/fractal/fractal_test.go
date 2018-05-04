package fractal

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"testing"
	"time"
)

const (
	_WIDTH = 400
	_HIGH  = 300
)

func Test_line_len(t *testing.T) {
	l := line_len(2, 1, 5, 5)
	if l != 5 {
		t.Errorf("The line len 5 != %d.", l)
	}
}

func Test_line(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, _WIDTH, _HIGH))

	cs := []color.RGBA{
		{255, 0, 0, 255},
		{255, 255, 0, 255},
		{0, 255, 0, 255},
		{0, 255, 255, 255},
		{0, 0, 255, 255},
		{255, 0, 255, 255},
	}
	x0 := _WIDTH / 2
	y0 := _HIGH / 2
	cs_i := 0
	for i := 0; i < _WIDTH; i += 7 {
		x := i
		y := 0
		cs_i++
		line(img, x, y, x0, y0, cs[cs_i%len(cs)])
	}
	for i := 0; i < _HIGH; i += 8 {
		x := _WIDTH
		y := i
		cs_i++
		line(img, x, y, x0, y0, cs[cs_i%len(cs)])
	}
	for i := _WIDTH; i >= 0; i -= 10 {
		x := i
		y := _HIGH
		cs_i++
		line(img, x, y, x0, y0, cs[cs_i%len(cs)])
	}
	for i := _HIGH; i >= 0; i -= 10 {
		x := 0
		y := i
		cs_i++
		line(img, x, y, x0, y0, cs[cs_i%len(cs)])
	}

	// Save to line.png
	f, _ := os.OpenFile("line.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func Test_koch(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 600, 600))
	// 等边三角形
	koch(img, 250, 0, 0, 433, color.RGBA{0, 0, 0, 255})
	koch(img, 0, 433, 500, 433, color.RGBA{0, 0, 0, 255})
	koch(img, 500, 433, 250, 0, color.RGBA{0, 0, 0, 255})
	f, _ := os.OpenFile("koch.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func Test_julia(t *testing.T) {
	const max_len = 1000
	img := image.NewRGBA(image.Rect(0, 0, max_len, max_len))
	julia(img, max_len)
	f, _ := os.OpenFile("julia.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func Test_mandelbrot(t *testing.T) {
	const max_len = 1000
	img := image.NewRGBA(image.Rect(0, 0, max_len, max_len))
	mandelbrot(img, max_len)
	f, _ := os.OpenFile("mandelbrot.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func Test_newton(t *testing.T) {
	const max_len = 500
	img := image.NewRGBA(image.Rect(0, 0, max_len, max_len))
	newton(img, max_len)
	f, _ := os.OpenFile("newton.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func Test_ifs(t *testing.T) {
	const (
		x = 300
		y = 400
	)
	img := image.NewRGBA(image.Rect(0, 0, x, y))
	ifs(img, x, y)
	f, _ := os.OpenFile("ifs.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func Test_triangle(t *testing.T) {
	const size = 500
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	A := image.Point{r.Intn(size), r.Intn(size)}
	B := image.Point{r.Intn(size), r.Intn(size)}
	C := image.Point{r.Intn(size), r.Intn(size)}
	fmt.Println(A, B, C)

	img1 := image.NewRGBA(image.Rect(0, 0, size, size))
	triangle_1(img1, B, A, C, color.Gray{0})
	f1, _ := os.OpenFile("triangle_1.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f1.Close()
	png.Encode(f1, img1)

	img2 := image.NewRGBA(image.Rect(0, 0, size, size))
	triangle_2(img2, B, A, C, color.Gray{0})
	f2, _ := os.OpenFile("triangle_2.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f2.Close()
	png.Encode(f2, img2)
}

func Test_sierpinski(t *testing.T) {
	const (
		x = 4000
		y = 3500
	)
	img := image.NewRGBA(image.Rect(0, 0, x, y))
	sierpinski(img, x, y)
	f, _ := os.OpenFile("sierpinski.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
