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
	nextEven := CreateEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
