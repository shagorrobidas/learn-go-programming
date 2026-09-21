package main
import "fmt"

func main() {
	// pointer or address of memory (ram)
	x := 10
	p := &x // p is a pointer to x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p)
}