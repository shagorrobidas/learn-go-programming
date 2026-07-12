package main

import "fmt"

var a = 10

func main() {
	fmt.Println("hello init function")
	fmt.Println(a)
}

func init() {
	fmt.Println("I am the first function that is executed first")
	fmt.Println(a)
	a = 12
}
