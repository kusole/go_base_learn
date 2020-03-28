package main

/**
n! = 1*2*3*...n
*/

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func fibonacci(x int) int {
	var retVal = 0
	if x == 1 {
		retVal = 1
	} else if x == 2 {
		retVal = 2
	} else {
		retVal = fibonacci(x-1) + fibonacci(x-2)
	}
	return retVal
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(fibonacci(5))
}
