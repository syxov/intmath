package intmath

import (
	"math"
	"math/rand"
	"testing"
)

func TestDim(t *testing.T) {
	if x := Dim(3, 5); x != 0 {
		t.Error("Expected", x, "to equal 0")
	}
	if x := Dim(5, 3); x != 2 {
		t.Error("Expected", x, "to equal 2")
	}
}

func TestAbs(t *testing.T) {
	if x := Abs(3); x != 3 {
		t.Error("Expected", x, "to equal 3")
	}
	if x := Abs(-3); x != 3 {
		t.Error("Expected", x, "to equal 3")
	}
}

func TestCopysign(t *testing.T) {
	if x := Copysign(3, 4); x != 3 {
		t.Error("Exprected", x, "to equal 3")
	}
	if x := Copysign(3, -4); x != -3 {
		t.Error("Exprected", x, "to equal -3")
	}
	if x := Copysign(-3, 4); x != 3 {
		t.Error("Exprected", x, "to equal 3")
	}
	if x := Copysign(-3, -4); x != -3 {
		t.Error("Exprected", x, "to equal -3")
	}
}

func BenchmarkMathCopysign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Copysign(float64(i), rand.Float64())
	}
}

func BenchmarkCopysign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copysign(i, rand.Int())
	}
}
