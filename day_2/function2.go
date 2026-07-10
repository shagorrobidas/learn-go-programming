package main

import "fmt"

func welcomeMessage(){
	fmt.Println("____Welcome to the Application____")
}

func byeMessage(){
	fmt.Println("Thanks for using the application")
	fmt.Println("Exiting the application")
}

func getNameAndAge() (string, int){
	var name string
	var age int
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Please enter your age: ")
	fmt.Scan(&age)
	return name, age
}

func printName(name string, age int)(string, int){
	return name, age
}

func getNumbers() (int, int){
	var num1, num2 int
	fmt.Print("Please enter your first number: ")
	fmt.Scan(&num1)
	fmt.Print("Please enter your second number: ")
	fmt.Scan(&num2)
	return num1, num2
}

func addNumbers(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func printResult(name string, age int, result int, num1 int, num2 int){
	fmt.Println("The name is:", name," and age is:", age)
	fmt.Println("The sum of", num1, "and", num2, "is:", result)
}

func main() {

	welcomeMessage()

	name, age := getNameAndAge()
	
	num1, num2 := getNumbers()

	name, age = printName(name, age)

	result := addNumbers(num1, num2)

	printResult(name, age, result, num1, num2)

	byeMessage()
}
