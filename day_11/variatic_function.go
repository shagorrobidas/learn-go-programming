package main 

import "fmt"


func printN(numbers ...int){
	
	fmt.Println("Slice", numbers)
	fmt.Println("Length", len(numbers))
	fmt.Println("Capecity", cap(numbers))

}
func main() {
	printN(5, 6, 7, 8, 9)

}