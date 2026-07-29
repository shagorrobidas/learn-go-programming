package main

import "fmt"

var x int = 10

func add(x int, y int) {
	z := x + y
	fmt.Println("The value of z is ", z)
}

func main() {
	add(5, 4)
	add(x, 10)

}
func init(){
	fmt.Println("Hello from init function")
}
