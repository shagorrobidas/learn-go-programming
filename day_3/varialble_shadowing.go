package main

import "fmt"

var (
	a = 10
	b = 20
)

func printNumber(num int) {
	fmt.Println("Number ", num)
}

func add(a int, b int) {
	res := a + b
	printNumber(res)

}

func main() {
	add(a, b)

}
