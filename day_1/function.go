package main

import "fmt"

func add (a int, b int) int {
	return a + b
}

func sub (a int, b int) int {
	return a - b
}

func mul (a int, b int) int {
	return a * b
}

func div (a int, b int) int {
	return a / b
}

func main() {
	var a int 
	var b int
	println("Enter first number : ")
	fmt.Scanln(&a)
	println("Enter second number : ")
	fmt.Scanln(&b)
	println("Sum is : ", add(a, b))
	println("Difference is : ", sub(a, b))
	println("Product is : ", mul(a, b))
	println("Quotient is : ", div(a, b))
}