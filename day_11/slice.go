package main
import "fmt"

func main() {


	// a := []int{10, 20, 30, 40, 50} // slice literal
	// fmt.Println("Slice elements:", a)
	// fmt.Println("Slice Length:", len(a))
	// fmt.Println("Slice Capacity:", cap(a))

	// arr := [6]string{"I", "Love", "Go", "Programming", "Language", "!"}
	// fmt.Println("Array elements:", arr)

	// s := arr[2:5]
	// fmt.Println("Slice elements:", s)
	// s1 := s[1:3]
	// fmt.Println("Slice elements s2 : ", s1)
	// fmt.Println("Slice Length:", len(s1))
	// fmt.Println("Slice Capacity:", cap(s1))


	// s := make([]int, 3)
	// s[0] = 10
	// s[2] = 20
	// fmt.Println("Slice elements:", s)
	// fmt.Println("Slice Length:", len(s))
	// fmt.Println("Slice Capacity:", cap(s))

	// a := make([]int, 3, 5)
	// a[0] = 10
	// a[1] = 20
	// a[2] = 30	
	// fmt.Println("Slice elements:", a)
	// fmt.Println("Slice Length:", len(a))
	// fmt.Println("Slice Capacity:", cap(a))
    
	var s []int // nil slice
	s = append(s, 10)
	s = append(s, 20)
	s = append(s, 30)
	s = append(s, 40)
	s = append(s, 50)
	fmt.Println("Slice elements:", s)
	fmt.Println("Slice Length:", len(s))
	fmt.Println("Slice Capacity:", cap(s))

}

/*
1. Slice from an existing array
2. Slice from a slice
3. slice literal
4. make function with len
5. make function with len and cap
6. empty or nil slice




*/