package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 1, 1
	arr := []int{15, 3, 1}
	result := 0
	output := 0.

	for _, v := range arr {
		result += v
	}

	for a < result {

		b = a + b
		a = b - a
	}

	output = (math.Abs(float64(a - result)))

	fmt.Println("Fibonacci:", a)
	fmt.Println("Result:", result)
	fmt.Println("Difference:", output)
}
