package main
import "fmt"


func changeSlice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}

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
    
	// var s []int // nil slice
	// s = append(s, 10)
	// s = append(s, 20)
	// s = append(s, 30)
	// s = append(s, 40)
	// s = append(s, 50)
	// s = append(s, 60, 70, 80, 90, 100)
	// fmt.Println("Slice elements:", s)
	// fmt.Println("Slice Length:", len(s))
	// fmt.Println("Slice Capacity:", cap(s))

	// a := s[2:5]
	// fmt.Println("Slice elements:", a)
	// fmt.Println("Slice Length:", len(a))
	// fmt.Println("Slice Capacity:", cap(a))
	// fmt.Println("Slice elements:", s[0:15])


	// var x []int
	// x = append(x, 1)
	// x = append(x, 2)
	// x = append(x, 3)
	
	// y:= x
	// x = append(x, 5)
	// y = append(y, 6)
	// y = append(y, 7)
	// fmt.Println("Slice elements x:", x)
	// fmt.Println("len(x):", len(x))
	// fmt.Println("cap(x):", cap(x))
	// fmt.Println("Slice elements y:", y)
	// fmt.Println("len(y):", len(y))
	// fmt.Println("cap(y):", cap(y))


	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]
	y := changeSlice(a)
	fmt.Println("x:", x[0:8])
	fmt.Println("y:", y)

	// fmt.Println("Slice elements x:", x)
	// fmt.Println("len(x):", len(x))
	// fmt.Println("cap(x):", cap(x))
	// fmt.Println("Slice elements y:", y)
	// fmt.Println("len(y):", len(y))
	// fmt.Println("cap(y):", cap(y))
}

/*
1. Slice from an existing array
2. Slice from a slice
3. slice literal
4. make function with len
5. make function with len and cap
6. empty or nil slice
7. slice underlying array rules => len and cap equal then append a new element to the slice .
then capacity will be doubled and new underlying array will be created and the existing elements 
will be copied to the new underlying array.
if len is less than cap then append a new element to the slice then the existing
underlying array will be used and the new element will be added to the existing underlying array.
rule =>less 1024 -> 100% increase
more than 1024 -> 25% increase





*/