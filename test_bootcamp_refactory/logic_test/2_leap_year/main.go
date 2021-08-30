package main

import "fmt"

func main() {
	start := 1900
	end := 2020

	fmt.Printf("Leap years between %v and %v:\n", start, end)
	for i := start + 1; i <= end; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}
