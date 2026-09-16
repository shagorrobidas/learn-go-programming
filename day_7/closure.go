package main

import "fmt"

const a = 10

var p = 1000

func outer() func() {
	money := 1000
	age := 20

	fmt.Println("Age is ", age)

	show := func() {
		money = money + a + p
		fmt.Println("Money is ", money)
	}

	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("Hello from init function")
}
