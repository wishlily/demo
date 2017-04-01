package fractal

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"
)

const (
	_IFS_PARAM_NUM = 7
)

func ifs_f(a, b, e, c, d, f float64, x, y float64) (float64, float64) {
	x1 := a*x + b*y + e
	y1 := c*x + d*y + f
	return x1, y1
}

func ifs_cal(a [][]float64, x, y float64, r *rand.Rand) (float64, float64) {
	if len(a) <= 0 || len(a[0]) != _IFS_PARAM_NUM {
		return x, y
	}
	p := r.Float64()
	c := 0.0
	for i := 0; i < len(a); i++ {
		c += a[i][_IFS_PARAM_NUM-1]
		if p <= c {
			x1, y1 := ifs_f(
				a[i][0], a[i][1], a[i][4],
				a[i][2], a[i][3], a[i][5],
				x, y)
			return x1, y1
		}
	}
	return x, y
}

func ifs_array_generate(r *rand.Rand) [][]float64 {
	a := [][]float64{}
	p := 0.0
	for p < 1 {
		tp := r.Float64()
		if p+tp <= 1 {
			p += tp
		} else {
			tp = 1 - p
			p = 1
		}
		arr := [_IFS_PARAM_NUM]float64{}
		arr[_IFS_PARAM_NUM-1] = tp
		for i := 0; i < _IFS_PARAM_NUM-1; i++ {
			arr[i] = 2*r.Float64() - 1
		}
		a = append(a, arr[:])
	}
	return a
}

func ifs(img *image.RGBA, limit_x, limit_y int) {
	a := [][]float64{
	// tree
	// {0.195, -0.488, 0.344, 0.433, 0.4431, 0.2452, 0.25},
	// {0.462, 0.414, -0.252, 0.361, 0.2511, 0.5692, 0.25},
	// {-0.058, -0.07, 0.453, -0.111, 0.5976, 0.0969, 0.25},
	// {-0.035, 0.07, -0.469, -0.022, 0.4884, 0.5069, 0.2},
	// {-0.637, 0, 0, 0.501, 0.8562, 0.2513, 0.05},
	// Sprial
	// {0.787879, -0.424242, 0.242424, 0.859848, 1.758647, 1.408065, 0.90},
	// {-0.121212, 0.257576, 0.151515, 0.053030, -6.721654, 1.377236, 0.05},
	// {0.181818, -0.136364, 0.090909, 0.181818, 6.086107, 1.568035, 0.05},
	// 蕨草叶
	// {0, 0, 0, 0.25, 0, -0.14, 0.02},
	// {0.85, 0.02, -0.02, 0.83, 0, 1, 0.84},
	// {0.09, -0.28, 0.3, 0.11, 0, 0.6, 0.07},
	// {-0.09, 0.28, 0.3, 0.09, 0, 0.7, 0.07},
	// 卷曲的叶子
	// {0, 0, 0, 0.25, 0, -0.04, 0.02},
	// {0.92, 0.05, -0.05, 0.93, -0.002, 0.5, 0.84},
	// {0.035, -0.2, 0.16, 0.04, -0.09, 0.02, 0.07},
	// {-0.04, 0.2, 0.16, 0.04, 0.083, 0.12, 0.07},
	// 茂盛的树
	// {-0.04, 0, -0.19, -0.47, -0.12, 0.3, 0.25},
	// {0.65, 0, 0, 0.56, 0.06, 1.56, 0.25},
	// {0.41, 0.46, -0.39, 0.61, 0.46, 0.4, 0.25},
	// {0.52, -0.35, 0.25, 0.74, -0.48, 0.38, 0.2},
	// 草
	// {0.786096, -0.704769, -0.4381057, -0.03372472, 0.142198, 0.799890, 0.181165},
	// {0.36127, -0.82840, -0.3979, -0.0091785, 0.097943, 0.173381, 0.33789},
	// {0.81895, 0.56741, 0.323338, 0.13015, -0.487087, -0.342847, 0.451090},
	// {0.99103, -0.534083, 0.47402, 0.1380517, -0.0976858, -0.384835, 0.0298530},
	// 螺旋
	// {-0.889272, 0.656922, -0.927109, -0.322378, 0.856565, 0.478706, 0.937723},
	// {0.411268, -0.581772, 0.998787, -0.631693, -0.731553, 0.226007, 0.0622774},
	// 羽球
	// {-0.872001, 0.0808141, 0.205013, 0.726129, -0.133041, -0.287383, 0.790841},
	// {0.148498, -0.880393, 0.901023, 0.567552, 0.99911, 0.760125, 0.209159},
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	if len(a) == 0 {
		a = ifs_array_generate(r)
		fmt.Printf("%.6v\n", a)
	}

	x0 := r.Float64()
	y0 := r.Float64()

	x, y := x0, y0
	x_min, x_max := x, x
	y_min, y_max := y, y

	loop := limit_x * limit_y / 2
	for i := 0; i < loop; i++ {
		x, y = ifs_cal(a, x, y, r)
		if x < x_min {
			x_min = x
		}
		if x > x_max {
			x_max = x
		}
		if y < y_min {
			y_min = y
		}
		if y > y_max {
			y_max = y
		}
	}
	x, y = x0, y0
	for i := 0; i < loop; i++ {
		x, y = ifs_cal(a, x, y, r)
		ix := (x - x_min) * float64(limit_x) / (x_max - x_min)
		iy := (y - y_min) * float64(limit_y) / (y_max - y_min)
		point(img, int(ix), int(iy), color.Gray{0})
	}
}