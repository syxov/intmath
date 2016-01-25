package intmath

import (
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
