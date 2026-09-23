package main
import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50} // slice literal
	fmt.Println("Slice elements:", a)
	fmt.Println("Slice Length:", len(a))
	fmt.Println("Slice Capacity:", cap(a))

	arr := [6]string{"I", "Love", "Go", "Programming", "Language", "!"}
	fmt.Println("Array elements:", arr)

	s := arr[2:5]
	fmt.Println("Slice elements:", s)
	s1 := s[1:3]
	fmt.Println("Slice elements s2 : ", s1)
	fmt.Println("Slice Length:", len(s1))
	fmt.Println("Slice Capacity:", cap(s1))
}

/*
1. Slice from an existing array
2. Slice from a slice
3. slice literal




*/