package main

import "fmt"

func main() {
	var x int
	var xPtr *int

	x = 10
	xPtr = &x

	fmt.Println(x)
	fmt.Println(xPtr)
	fmt.Println(*xPtr)
	fmt.Println(&xPtr)
	fmt.Println(*&xPtr)
}
