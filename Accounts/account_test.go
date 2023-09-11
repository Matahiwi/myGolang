package Accounting

import "testing"

func TestTotal(t *testing.T) {
	y := []int{1, 2, 3, 4, 5}
	x := Total(y)
	if x != 15 {
		t.Error("Expected", 15, "got", x)
	}
}

func TestAverage(t *testing.T) {
	y := []int{1, 2, 3, 4, 5}
	x := Average(y)
	if x != 3 {
		t.Error("Expected", 3, "got", x)
	}
}
