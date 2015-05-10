package intmath

import (
	"testing"
)

func TestSign(t *testing.T) {
	if x := Sign(-10); x != -1 {
		t.Error("Expected", x, "to equal -1")
	}
	if x := Sign(10); x != 1 {
		t.Error("Expected", x, "to equal 1")
	}
	if x := Sign(0); x != 0 {
		t.Error("Expected", x, "to equal 0")
	}
}

func TestAverage(t *testing.T) {
	if x := Average(3, 5); x != 4 {
		t.Error("Expected", x, "to equal 4")
	}
}

func TestDim(t *testing.T) {
	if x := Dim(3, 5); x != 0 {
		t.Error("Expected", x, "to equal 0")
	}
	if x := Dim(5, 3); x != 2 {
		t.Error("Expected", x, "to equal 2")
	}
}
