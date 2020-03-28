package main

import "fmt"

func change(x *int) {
	*x = 200
}

func setValue(xPtr *int) {
	*xPtr =100
}

func main(){
	var x int = 100
	fmt.Println(x)
	change(&x)
	fmt.Println(x)

	xPtr :=new (int)
	setValue(xPtr)
	fmt.Println(xPtr)
	fmt.Println(&xPtr)
	fmt.Println(*xPtr)
}