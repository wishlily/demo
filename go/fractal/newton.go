package fractal

import (
	"image"
	"image/color"
	"math/cmplx"
)

const (
	_NEWTON_E       = 0.001
	_NEWTON_MAP_MAX = 3.0
)

func newton_f(z complex128) complex128 {
	// f := z*z + 1 // z^2+1
	// f := z*z*z - 1 // z^3-1
	// f := z*z*z*z - 1 // z^4-1
	// f := z*z*z*z*z - 1 // z^5-1
	// f := cmplx.Pow(z, 8) + 15*cmplx.Pow(z, 4) - 16 // z^8+15x^4-16
	// f := cmplx.Pow(z, 4+3i) - 1 // z^(4+3i)-1
	f := cmplx.Pow(z, 7) + 1i*z - 1
	return f
}

func newton_df(z complex128) complex128 {
	// df := 2 * z // z^2+1
	// df := 3 * z * z // z^3-1
	// df := 4 * z * z * z // z^4-1
	// df := 5 * z * z * z * z // z^5-1
	// df := 8*cmplx.Pow(z, 7) + 60*cmplx.Pow(z, 3) // z^8+15x^4-16
	// df := (4 + 3i) * cmplx.Pow(z, 3+3i) // z^(4+3i)-1
	df := 7*cmplx.Pow(z, 6) + 1i
	return df
}

func newton_iter(z complex128) int {
	for i := 255; i >= 0; i -= 9 {
		z1 := z - newton_f(z)/newton_df(z)
		if cmplx.Abs(z1-z) < _NEWTON_E {
			return i
		}
		z = z1
	}
	return 0
}

func newton(img *image.RGBA, limit int) {
	for i := 0; i < limit; i++ {
		zx := _NEWTON_MAP_MAX*float64(i)/float64(limit) - _NEWTON_MAP_MAX/2
		for j := 0; j < limit; j++ {
			zy := _NEWTON_MAP_MAX*float64(j)/float64(limit) - _NEWTON_MAP_MAX/2
			gray := uint8(newton_iter(complex(zx, zy)))
			point(img, i, j, color.Gray{gray})
		}
	}
}
