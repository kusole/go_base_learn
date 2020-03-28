package main

import (
	"fmt"
)

type Rect struct {
	width  float64
	length float64
}

func doubleArea(rect Rect) float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}

func main() {
	var rect = Rect{
		width:  100,
		length: 200,
	}
	fmt.Println(doubleArea(rect))
	fmt.Println("Width:", rect.width, "Length:", rect.length)
}
