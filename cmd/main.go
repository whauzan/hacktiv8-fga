package main

import "fmt"

func main() {
	fmt.Println("Hello! Wahyu here ^-^")

	var i int
	var j bool
	var k float64 = 123.456

	i = 21

	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")

	j = true

	fmt.Printf("%t \n", j)

	j = false

	fmt.Printf("%t \n", j)
	fmt.Printf("%+q \n", "Я")
	fmt.Printf("%d \n", 21)
	fmt.Printf("%o \n", 25)
	fmt.Printf("%x \n", "f")
	fmt.Printf("%X \n", "F")
	fmt.Printf("\u042F \n")
	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)
}