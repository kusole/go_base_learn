package main

import (
	"fmt"
)

// Rect is decriber
type Rect struct {
	width  float64
	length float64
}

/*
type Rect struct {
	width, length float64
}
*/

func main() {
	var rect = Rect{
		width:  200,
		length: 100,
	}
	/*
		var rect = Rect {width: 200, length: 100}
	*/
	fmt.Println(rect.length)
}
