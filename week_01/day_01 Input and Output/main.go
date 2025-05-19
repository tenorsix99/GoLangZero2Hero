package main

import "fmt"

func main() {
	var name string = "GoLang"
	age := 10
	price := 3.99
	isCool := true

	fmt.Println("ชื่อ:", name)
	fmt.Println("อายุ:", age)
	fmt.Println("ราคา:", price)
	fmt.Println("เท่ไหม:", isCool)

	var a int
	var b int

	fmt.Println("ค่า a =")
	fmt.Scanln(&a)
	fmt.Println("ค่า b =")
	fmt.Scanln(&b)

	fmt.Printf("%d + %d = %d \n", a, b, (a + b))
	fmt.Printf("%d - %d = %d \n", a, b, (a - b))
	fmt.Printf("%d * %d = %d \n", a, b, (a * b))
	fmt.Printf("%d / %d = %d \n", a, b, (a / b))
}
