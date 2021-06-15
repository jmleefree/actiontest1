package main

import (
	lib "actiontest1/lib"
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSum3(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum3 was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSum4(t *testing.T) {
	total := lib.SumLib2(5, 5)
	if total != 10 {
		t.Errorf("Sum4 was incorrect, got: %d, want: %d.", total, 10)
	}
}
