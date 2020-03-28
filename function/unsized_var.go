package main

import (
	"fmt"
)

func sum(arr ...int) int {
	sum := 0
	for _, item := range arr {
		sum += item
	}
	return sum
}

func sum1(base int, arr ...int) int {
	sum := base
	for _, item := range arr {
		sum += item
	}
	return sum
}

func main() {
	// 自带的变长
	fmt.Println(1)
	fmt.Println(1, 2)
	fmt.Println(1, 2, 3)

	// 变长参数传参数
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	var arr = []int{1, 2, 3}
	fmt.Println(arr)

	// 变长参数只能是最后一个
	fmt.Println(sum1(100, 1, 2, 3))

	// 变长 + 数组
	fmt.Println(sum1(100, arr...))
}
