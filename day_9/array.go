package main 

import "fmt"


func main(){
	var n int
	fmt.Print("Enter the size: ")
	fmt.Scanln(&n)

	// Array lengths in Go must be constants.
	// Use make() to create a slice with a dynamic size at runtime:
	arr := make([]int, n)

	fmt.Println("Slice:", arr)
}