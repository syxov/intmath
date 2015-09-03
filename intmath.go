package intmath

var (
	intSizeMinusOne uint
)

func init() {
	intSizeMinusOne = (32 << (^uint(0) >> 63)) - 1
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

func Sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func Copysign(x, y int) int {
	t := (x ^ y) >> intSizeMinusOne
	return (x ^ t) - t
}

func Abs(x int) int {
	y := x >> intSizeMinusOne
	return (x ^ y) - y
}

func Dim32(x, y int32) (value int32) {
	if x > y {
		value = x - y
	}
	return
}

func Max32(x, y int32) int32 {
	if x >= y {
		return x
	}
	return y
}

func Min32(x, y int32) int32 {
	if x <= y {
		return x
	}
	return y
}

func Average32(x, y int32) int32 {
	return (x & y) + ((x ^ y) >> 1)
}

func Signbit32(x int32) bool {
	return x < 0
}

func Sign32(x int32) int32 {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func Copysign32(x, y int32) int32 {
	t := (x ^ y) >> 31
	return (x ^ t) - t
}

func Abs32(x int32) int32 {
	y := x >> 31
	return (x ^ y) - y
}
