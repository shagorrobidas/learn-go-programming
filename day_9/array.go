package main 

import "fmt"

var arr2 = [3] string{"I","Love","Go"}


func main(){
	
	// var arr [2] int

	// arr[1] = 10
	// arr[0] = 20
	// fmt.Println(arr)

	// arr1 := [3] int{10,20,30}
	// fmt.Println(arr1)
	// fmt.Println(arr2)
	// fmt.Println(arr2[1])

	var n int
	fmt.Print("Enter the size: ")
	fmt.Scanln(&n)

	// Create a slice with dynamic size at runtime:
	arr := make([]int, n)

	// Read number inputs into the slice
	fmt.Println("Enter", n, "numbers:")
	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scanln(&arr[i])
	}

	fmt.Println("Your Slice/Array:", arr)
}