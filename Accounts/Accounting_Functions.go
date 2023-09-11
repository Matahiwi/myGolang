package Accounting

func Total(x []int) int {
	total := 0
	for _, val := range x {
		total += val
	}
	return total
}

func Average(x []int) int {
	total := 0
	for _, val := range x {
		total += val
	}
	average := total / len(x)
	return average
}
