package main

import (
	"Accounting"
	"Mathematics"
	"fmt"
)

func main() {
	// added a comment
	x := []int{1, 2, 3, 4, 5}
	fmt.Println("Total:", x, Accounting.Total(x))
	fmt.Println("Average:", x, Accounting.Average(x))
	for i := 5; i >= 0; i-- {
		fmt.Println(i, Mathematics.Factorial(i), Mathematics.Addtorial(i))
	}
}
