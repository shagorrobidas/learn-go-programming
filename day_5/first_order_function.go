package main
import "fmt"





func processOpration(a int , b int, opration func(x int,y int)) {
	opration(a,b)
}


func add(a int, b int) {
	c := a + b
	fmt.Println("The sum is ", c)

}

func sub(a int, b int) {
	c := a - b
	fmt.Println("The sub is ", c)

}
func mul(a int, b int) {
	c := a * b
	fmt.Println("The mul is ", c)

}
func div(a int, b int) {
	c := a / b
	fmt.Println("The div is ", c)

}

func mod(a int, b int) {
	c := a % b
	fmt.Println("The mod is ", c)

}

func switchOpration(a int, b int, opration int) {
	switch opration {
	case 1:
		add(a, b)
	case 2:
		sub(a, b)
	case 3:
		mul(a, b)
	case 4:
		div(a, b)
	case 5:
		mod(a, b)
	default:
		fmt.Println("Invalid opration")
	}
}
func call()func (a int, b int){
	return sub
}


func main() {
	var a , b int 
	fmt.Printf("Enter first number :\t")
	fmt.Scan(&a)
	fmt.Printf("Enter second number :\t")
	fmt.Scan(&b)

	addFunc := call()  // here i am getting a function exprassion
	addFunc(a,b) // here i am calling the function
	
	// processOpration(a,b , add)
	// processOpration(a,b , sub)
	// switchOpration(a,b , 1)
	// switchOpration(a,b , 2)
	// switchOpration(a,b , 3) 
	// switchOpration(a,b , 4)
	// switchOpration(a,b , 5)
}




/*

	1. perameter vs argument
	2. First order function
		i. Standard function or named function
		ii. Anonymous function
		iii.IIFE function
		iv. function expression

	3. Higher order function


	functional paradigm -> haskel,  racket
	 math -> logic (discrete mathematics)

	 1. first order logic
	 2. higher order logic


	 #### LOGIC  ####

	 1. Object (pepole, animal, plants, place, thing) = > Subject
	 2. Property (Color, Size, Shape, Texture, etc) => Predicate
	 3. Relation


*/
