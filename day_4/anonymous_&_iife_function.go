package main

import "fmt"

// anonymous function
func main() {
	var num1, num2 int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)
    // anonymous function
	// Immediately Invoked Function Expression (IIFE)
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(num1, num2)


}

func init(){
	fmt.Println("I'll be called first")
}