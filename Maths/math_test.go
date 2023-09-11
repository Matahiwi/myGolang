package Mathematics

import "testing"

func TestFactorial(t *testing.T) {
	x := Factorial(4)
	//if x != 24 {
	if x != 24 {
		t.Error("Expected:", 24, "Got", x)
	}
}

func TestAddtorial(t *testing.T) {
	x := Addtorial(4)
	if x != 10 {
		t.Error("Expected:", 10, "Got", x)
	}
}
