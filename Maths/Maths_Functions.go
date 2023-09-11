package Mathematics

func Factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}

func Addtorial(x int) int {
	if x == 0 {
		return 0
	}
	return x + Addtorial(x-1)
}
