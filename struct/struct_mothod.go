package main

import "fmt"

type Rect struct {
	width  float64
	length float64
}

func (rect Rect) area() float64 {
	return rect.width * rect.length
}

func main() {
	var rect = Rect{100, 200}
	fmt.Println(rect.area())
}
