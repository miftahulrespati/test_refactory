package main

import (
	"fmt"
	"strconv"
)

func main() {
	const n = 15
	arr := [n]string{}
	var str string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			str = "FizzBuzz"
			arr[i-1] = str
		} else if i%3 == 0 {
			str = "Fizz"
			arr[i-1] = str
		} else if i%5 == 0 {
			str = "Buzz"
			arr[i-1] = str
		} else {
			str = strconv.Itoa(i)
			arr[i-1] = str
		}
	}
	fmt.Println(arr)
}
