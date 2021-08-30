package main

import "fmt"

func main() {
	a, b := 1, 1
	arr := []int{15, 3, 1}
	result := 0
	output := 0

	for _, v := range arr {
		result += v
	}

	for a < result-1 {

		b = a + b
		a = b - a
	}

	if a < result {

		output = (result - a)
	} else if result < a {
		output = (a - result)
	}

	fmt.Println("Fibonacci:", a)
	fmt.Println("Result:", result)
	fmt.Println("Selisih:", output)
}
