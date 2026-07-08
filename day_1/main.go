package main

import (
	"fmt"
)

func main() {
	// Print "Hello World!" to the console
	/*	This is a multi-line comment.
		It can span multiple lines.
		Useful for longer explanations or documentation.
	*/
	// fmt.Println("Hello World!") 
	// fmt.Println("This line does not execute")

	// variable declaration and initialization
	var name string = "Shagor Robidas"
	var age int = 26
	var address string = "Dhaka, Bangladesh"
	var isVerified bool = false
	fmt.Println("My name is", name + " and I am", age, "years old.", "I live in", address + ".")
	fmt.Println("Is verified:", isVerified)

	// var a string = "" // This is an empty string, which is the default value for string variables
	// println("Value of a before initialization:", a) // This will print an empty string
	// a = "Now I have a value!"
	// println("Value of a after initialization:", a) // This will print the assigned value

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var name1 string 
	name1 = "Shagor Robidas"
	fmt.Println(name)
}