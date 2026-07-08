package main

import "fmt"


func main() {
	// >, <, >=, <=, ==, !=
	// and = &&
	// or = ||
	// not = !

	var age int

	fmt.Print("Enter your age : ")
	fmt.Scanln(&age)


	// if age > 18 {
	// 	println("You are eligible to vote")
	// } else if age < 18 {
	// 	println("You are not eligible to vote")
	// } else if age == 18 {
	// 	println("You are eligible to vote")
	// }else {
	// 	println("You are not eligible to vote")
	// }

	// if age ==20 && sex == "male" {
	// 	println("You are ready to be married")
	// }

	switch age {
	case 18:
		println("You are eligible to vote")
	case 20:
		println("You are ready to be married")
	default:
		println("You are not eligible to vote")
	}



}