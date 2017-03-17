package fractal

import (
	"image"
	"image/color"
)

const (
	_KOCH_LINE_MIN_LENGTH = 15
)

// line -> 5 point
func koch_point(x1, y1, x2, y2 int) []image.Point {
	var pots [5]image.Point

	pots[0].X = x1
	pots[0].Y = y1
	pots[4].X = x2
	pots[4].Y = y2

	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	pots[1].X = x1 + int(dx/3)
	pots[1].Y = y1 + int(dy/3)
	pots[3].X = x2 - int(dx/3)
	pots[3].Y = y2 - int(dy/3)

	/* http://blog.csdn.net/u013445530/article/details/44904017
	 * 逆时针
	 */
	dx = float64(pots[3].X - pots[1].X)
	dy = float64(pots[3].Y - pots[1].Y)
	pots[2].X = int(dx*0.5-dy*0.8660254) + pots[1].X
	pots[2].Y = int(dy*0.5+dx*0.8660254) + pots[1].Y

	return pots[:]
}

/*
 * 维基百科: 科赫曲线
 * 给定线段AB，科赫曲线可以由以下步骤生成：
 * 1. 将线段分成三等份（AC,CD,DB）
 * 2. 以CD为底，向外（内外随意）画一个等边三角形DMC
 * 3. 将线段CD移去
 * 4. 分别对AC,CM,MD,DB重复1~3。
 */
func koch(img *image.RGBA, x1, y1, x2, y2 int, c color.Color) {
	p := koch_point(x1, y1, x2, y2)
	next := true
	for i := 0; i < len(p)-1; i++ {
		if line_len(p[i].X, p[i].Y, p[i+1].X, p[i+1].Y) < _KOCH_LINE_MIN_LENGTH {
			next = false
		}
	}
	for i := 0; i < len(p)-1; i++ {
		if next {
			koch(img, p[i].X, p[i].Y, p[i+1].X, p[i+1].Y, c)
		} else {
			line(img, p[i].X, p[i].Y, p[i+1].X, p[i+1].Y, c)
		}
	}
}
