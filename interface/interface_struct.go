package main

import "fmt"

type Phone interface {
	sales() int
}

type NokiaPhone struct {
	price int
}

func (phone NokiaPhone) sales() int {
	return phone.price
}

type IPhone struct {
	price int
}

func (phone IPhone) sales() int {
	return phone.price
}

type Person struct {
	phones []Phone
	name string
	age int
}

func (person Person) total_cost() int {
	var sum = 0
	for _, phone:=range person.phones {
		sum += phone.sales()
	}
	return sum
}

func main() {
	var bought_phones = [5]Phone {
		NokiaPhone{price: 350},
        IPhone{price: 5000},
        IPhone{price: 3400},
        NokiaPhone{price: 450},
        IPhone{price: 5000},
	}	
	
	var person = Person{phones: bought_phones[:], name: "liutao", age: 18}

	fmt.Println(person.phones[0].sales())
	
	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.total_cost())
}