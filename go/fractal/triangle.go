package fractal

import (
	"image"
	"image/color"
)

const (
	_TRIANGLE_SIDE = false
)

func triangle_cross(a, b, c image.Point) int {
	ab := image.Point{b.X - a.X, b.Y - a.Y}
	ac := image.Point{c.X - a.X, c.Y - a.Y}
	return ab.X*ac.Y - ab.Y*ac.X
}

func triangle_right_side(a, b, c image.Point) bool {
	if triangle_cross(a, b, c) > 0 {
		return true
	}
	return false
}

/*
 * 排序
 * 最靠近 x 轴的点为第一个点，剩下逆时针排序
 */
func triangle_sort(a, b, c image.Point) (image.Point, image.Point, image.Point) {
	sort := []image.Point{a, b, c}
	n := 0
	for i := 1; i < 3; i++ {
		if sort[i].Y < sort[n].Y {
			n = i
		} else if sort[i].Y == sort[n].Y {
			if sort[i].X < sort[n].X {
				n = i
			}
		}
	}
	if n != 0 {
		t := sort[0]
		sort[0] = sort[n]
		sort[n] = t
	}
	if triangle_right_side(sort[0], sort[1], sort[2]) == _TRIANGLE_SIDE {
		return sort[0], sort[1], sort[2]
	}
	return sort[0], sort[2], sort[1]
}

func triangle_min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func triangle_max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func triangle_1(img *image.RGBA, A, B, C image.Point, c color.Color) {
	A, B, C = triangle_sort(A, B, C)
	// fmt.Println(A, B, C)

	// 扫描正方形
	x_min, x_max := A.X, A.X
	y_min, y_max := A.Y, A.Y

	x_min = triangle_min(x_min, B.X)
	x_min = triangle_min(x_min, C.X)
	x_max = triangle_max(x_max, B.X)
	x_max = triangle_max(x_max, C.X)

	y_min = triangle_min(y_min, B.Y)
	y_min = triangle_min(y_min, C.Y)
	y_max = triangle_max(y_max, B.Y)
	y_max = triangle_max(y_max, C.Y)

	for i := x_min; i <= x_max; i++ {
		for j := y_min; j <= y_max; j++ {
			D := image.Point{i, j}
			if triangle_right_side(A, B, D) == _TRIANGLE_SIDE &&
				triangle_right_side(B, C, D) == _TRIANGLE_SIDE &&
				triangle_right_side(C, A, D) == _TRIANGLE_SIDE {
				img.Set(i, j, c)
			}
		}
	}
}

/* BUG: 一些锐角三角形无法绘出 */
func triangle_2(img *image.RGBA, A, B, C image.Point, c color.Color) {
	A, B, C = triangle_sort(A, B, C)
	// fmt.Println(A, B, C)

	// 蛇字形扫描
	x_min, x_max := A.X, A.X
	x_min = triangle_min(x_min, B.X)
	x_min = triangle_min(x_min, C.X)
	x_max = triangle_max(x_max, B.X)
	x_max = triangle_max(x_max, C.X)

	x := A.X
	x_step := 1
	for y, loop := A.Y, true; loop; y++ {
		loop = false
		for ; x >= x_min && x <= x_max; x += x_step {
			D := image.Point{x, y}
			if triangle_right_side(A, B, D) == _TRIANGLE_SIDE &&
				triangle_right_side(B, C, D) == _TRIANGLE_SIDE &&
				triangle_right_side(C, A, D) == _TRIANGLE_SIDE {
				loop = true
				img.Set(x, y, c)
			} else if loop {
				break
			}
		}
		x_step *= -1
	}
}
