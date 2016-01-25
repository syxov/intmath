package intmath

import (
	"math"
)

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

func Acos(x int) float64 {
	return math.Acos(float64(x))
}

func Acosh(x int) float64 {
	return math.Acosh(float64(x))
}

func Asin(x int) float64 {
	return math.Asin(float64(x))
}

func Asinh(x int) float64 {
	return math.Asinh(float64(x))
}

func Atan(x int) float64 {
	return math.Atan(float64(x))
}

func Atan2(y, x int) float64 {
	return math.Atan2(float64(y), float64(x))
}

func Atanh(x int) float64 {
	return math.Atanh(float64(x))
}

func Cbrt(x int) float64 {
	return math.Cbrt(float64(x))
}

func Copysign(x, y int) int {
	return int(math.Copysign(float64(x), float64(y)))
}

func Cos(x int) float64 {
	return math.Cos(float64(x))
}

func Cosh(x int) float64 {
	return math.Cosh(float64(x))
}

func Dim(x, y int) int {
	if x > y {
		return x - y
	}
	return 0
}

func Erf(x int) float64 {
	return math.Erf(float64(x))
}

func Erfc(x int) float64 {
	return math.Erfc(float64(x))
}

func Exp(x int) float64 {
	return math.Exp(float64(x))
}

func Exp2(x int) float64 {
	return math.Exp2(float64(x))
}

func Expm1(x int) float64 {
	return math.Expm1(float64(x))
}

func Gamma(x int) float64 {
	return math.Gamma(float64(x))
}

func Hypot(p, q int) float64 {
	return math.Hypot(float64(p), float64(q))
}

func Ilogb(x int) int {
	return math.Ilogb(float64(x))
}

func J0(x int) float64 {
	return math.J0(float64(x))
}

func J1(x int) float64 {
	return math.J1(float64(x))
}

func Jn(n int, x int) float64 {
	return math.Jn(n, float64(x))
}

func Ldexp(frac int, exp int) float64 {
	return math.Ldexp(float64(frac), exp)
}

func Lgamma(x int) (lgamma float64, sign int) {
	return math.Lgamma(float64(x))
}

func Log(x int) float64 {
	return math.Log(float64(x))
}

func Log10(x int) float64 {
	return math.Log10(float64(x))
}

func Log1p(x int) float64 {
	return math.Log1p(float64(x))
}

func Log2(x int) float64 {
	return math.Log2(float64(x))
}

func Logb(x int) float64 {
	return math.Logb(float64(x))
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Sin(x int) float64 {
	return math.Sin(float64(x))
}

func Sincos(x int) (sin, cos float64) {
	return math.Sincos(float64(x))
}

func Sinh(x int) float64 {
	return math.Sinh(float64(x))
}

func Sqrt(x int) float64 {
	return math.Sqrt(float64(x))
}

func Tan(x int) float64 {
	return math.Tan(float64(x))
}

func Tanh(x int) float64 {
	return math.Tanh(float64(x))
}

func Y0(x int) float64 {
	return math.Y0(float64(x))
}

func Y1(x int) float64 {
	return math.Y1(float64(x))
}

func Yn(n int, x int) float64 {
	return math.Yn(n, float64(x))
}
