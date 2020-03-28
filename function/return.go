package main

import (
	"fmt"
)

// name return
func sliceSum(arr []int) (sum int) {
	sum = 0
	for _, elem := range arr {
		sum += elem
	}
	return
}

// mutli retrun
func sliceSum1(arr []int) (int, float64) {
	sum := 0
	avg := 0.0
	for _, elem := range arr {
		sum += elem
	}
	avg = float64(sum) / float64(len(arr))
	return sum, avg
}

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(sliceSum(arr))
	fmt.Println(sliceSum1(arr))
}
