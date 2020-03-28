package main

import (
	"fmt"
)

type Phone struct {
	price int
	color string
}

func (phone Phone) ringring() {
	fmt.Println("Phone is ringing...")
}

type IPhone struct {
	Phone
	model string
}

func main() {
	var p IPhone
	p.price = 5000
	p.color = "black"
	p.model = "iphone 5"
	fmt.Println("i have a iphone")
	fmt.Println("price: ", p.price)
	fmt.Println("color: ", p.color)
	fmt.Println("model: ", p.model)

	p.ringring()
}
