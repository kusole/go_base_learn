package main

import (
	"fmt"
)

type Rect struct {
	width  float64
	length float64
}

func (rect *Rect) area() float64 {
	return rect.length * rect.width
}

func (rect *Rect) area_double() float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.length * rect.width
}

func main() {
	var rect = new(Rect)
	rect.width = 100
	rect.length = 200
	fmt.Println(*rect)
	fmt.Println(rect.area())
	fmt.Println(*rect)
	fmt.Println(rect.area_double())
	fmt.Println(*rect)
}
