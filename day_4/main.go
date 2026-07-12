package main
import "fmt"

var a = 10

func main(){
	age := 18
	if age >= 18 {
		a := 20
		fmt.Println("a value : ", a)
	} else {
		fmt.Println("You are not eligible for voting")
	}


	fmt.Println("a value : ", a)

}
