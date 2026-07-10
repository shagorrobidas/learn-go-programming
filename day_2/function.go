package main

import "fmt"

func add(a int, b int) int{
	sum := a + b
	return sum
}

func getNumber(num1 int, num2 int) (int, int, int, int) {
	sum := num1 + num2

	mul := num1 * num2

	sub := num1 - num2

	div := num1 / num2

	return sum, mul, sub, div
}

func printSomthing(){
	fmt.Println("This is a function")
}

func sayHello(name string){
	fmt.Println("Hello", name)
}

func main() {
	// a := 10
	// b := 20



	// result := add(a, b)
	// fmt.Println("The sum of", a, "and", b, "is:", result)

	// sum, mul, sub, div := getNumber(a, b)
	// fmt.Println("The sum of", a, "and", b, "is:", sum)
	// fmt.Println("The mul of", a, "and", b, "is:", mul)
	// fmt.Println("The sub of", a, "and", b, "is:", sub)
	// fmt.Println("The div of", a, "and", b, "is:", div)

	printSomthing()

	var name string

	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)

	sayHello(name)

}