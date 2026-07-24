package main

import "fmt"

var sum = func(a int, b int) {
	fmt.Println(a + b)
}

func init() {
	fmt.Println("I'll be called first")
}

func main() {
	var num1, num2 int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	sum(num1, num2)
}
