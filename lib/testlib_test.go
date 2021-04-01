package lib

import "testing"

func TestSumLib1(t *testing.T) {
	total := SumLib1(5, 5)
	if total != 10 {
		t.Errorf("TestSumLib1 was incorrect, got: %d, want: %d.", total, 10)
	}
}
