package main

import "fmt"

type Phone interface {
	call()
	sales() int
}

type NokiaPhone struct {
	price int
}

func (nokia NokiaPhone) call() {
	fmt.Println("i am nokia, i can call you")
}

func (nokia NokiaPhone) sales() int {
	return nokia.price
}

type IPhone struct {
	price int
}

func (iphone IPhone) call() {
	fmt.Println("i an iphone, i can call you")
}

func (iphone IPhone) sales() int {
	return iphone.price
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	// phone.price = 5000
	phone.call()
	fmt.Println(phone.sales())

	phone = new(IPhone)
	// phone.price = 6000
	phone.call()
	fmt.Println(phone.sales())

	var phones = [5]Phone{
		NokiaPhone{price: 350},
		IPhone{price: 5000},
		IPhone{price: 3400},
		NokiaPhone{price: 450},
		IPhone{price: 5000},
	}

	var totalSales = 0
	for _, phone := range phones {
		totalSales += phone.sales()
	}
	fmt.Println(totalSales)
}
