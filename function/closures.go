package main

import "fmt"

// CreateEvenGenerator is a function for func
func CreateEvenGenerator() func() uint {
	i := uint(0)
	return func() (retVal uint) {
		retVal = i
		i += 2
		return
	}
}

func main() {
	var arr1 = []int{1, 2, 3, 4}
	var sum = func(arr ...int) int {
		totalSum := 0
		for _, val := range arr {
			totalSum += val
		}
		return totalSum
	}
	fmt.Println(sum(arr1...))

	nextEven := CreateEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
