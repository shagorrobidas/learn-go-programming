package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 20

	fmt.Println("Age is ", age)

	show := func() {
		money = money + a + p
		fmt.Println("Money is ", money)
	}
	// scape analysis: The variables 'money' and 'age' are captured by the closure 'show' and will be allocated on the heap.	

	return show
}

func call() {
	incr1 := outer()  // show is a closure that captures the variables 'money' and 'age' from the outer function 'outer'. The closure is returned to the caller and can be invoked later. memory and age are allocated on the heap because they are captured by the closure. The closure can be invoked multiple times and will have access to the same variables 'money' and 'age' that were captured when the closure was created.
	incr1()     // show = money + a + p = 100 + 10 + 100 = 210
	incr1()   // show = money + a + p = 210 + 10 + 100 = 320

	incr2 := outer()
	incr2()   // show = money + a + p = 100 + 10 + 100 = 210
	incr2()   // show = money + a + p = 210 + 10 + 100 = 320
	incr2()   // show = money + a + p = 320 + 10 + 100 = 430
}

func main() {
	call()
}

func init() {
	fmt.Println("======== bank  ===============")
}
