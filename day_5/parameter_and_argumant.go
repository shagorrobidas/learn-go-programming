package main

import "fmt"

func add(a int, b int){   // parameter => a,b
	c := a+b
	fmt.Println("The Sum is ", c)
}

func main(){
	fmt.Println("Hello from main function")
	
	// call the add function
	add(10, 20) // argument => 10, 20
}


func init(){
	fmt.Println("Hello from init function")
}