package main
import "fmt"


// pass by value 
// pass by reference (pointer)

func printArray(numbers *[3]int) {
	fmt.Println("Array elements:", numbers)


}

func printStrArray(words *[3]string) {
	fmt.Println("Array elements:", *words)
	fmt.Println("Address of array:", words)
}

type User struct {
	Name     string
	Age      int
	Slary    float64
	FavFoods []string
}

func main() {
	// pointer or address of memory (ram)
	// x := 10
	// fmt.Println("Value of x:", x)
	// p := &x // p is a pointer to x
	// *p = 20 // dereferencing p to modify the value of x
	// fmt.Println("Value of x:", x)
	// fmt.Println("Address of x:", p) // p is the address of x
	// fmt.Println("Value at address p:", *p) // dereferencing p to get the value of x

	// arr := [3]int{10, 20, 30}
	// fmt.Println("Array before function call:", arr)
	// printArray(&arr)

	// arr1 := [3]string{"I", "Love", "Go"}
	// fmt.Println("Array before function call:", arr1)
	// printStrArray(&arr1)

	user1 := User{ // instance or object of struct
		Name:     "John",
		Age:      30,
		Slary:    50000.0,
		FavFoods: []string{"Burger", "Pizza"},
	}
	fmt.Println("User details:", user1)
	fmt.Println("User name "+ user1.Name + ", Age: ", user1.Age, ", Salary: ", user1.Slary)

	p := &user1 // pointer to user1
	fmt.Println("Address of user1:", p)
	fmt.Println("Value at address p:", *p)
	fmt.Println("User name "+ p.Name + ", Age: ", p.Age, ", Salary: ", p.Slary)


}