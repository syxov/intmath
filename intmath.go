package intmath

var (
	intSizeMinusOne, bitField uint
)

func init() {
	intSizeMinusOne = (32 << (^uint(0) >> 63)) - 1
	bitField = 1 << intSizeMinusOne
}

func Dim(x, y int) (value int) {
	if x > y {
		value = x - y
	}
	return
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func Average(x, y int) int {
	return (x & y) + ((x ^ y) >> 1)
}

func Signbit(x int) bool {
	return x < 0
}

func Sign(x int) (result int) {
	if x > 0 {
		result = 1
	} else if x < 0 {
		result = -1
	}
	return
}

func Copysign(x, y int) int {
	t := (x ^ y) >> intSizeMinusOne
	return (x ^ t) - t
}

func Abs(x int) int {
	y := x >> intSizeMinusOne
	return (x ^ y) - y
}
