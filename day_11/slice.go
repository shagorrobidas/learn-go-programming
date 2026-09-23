package main
import "fmt"

func main() {
	arr := [6] string{"I", "Love", "Go", "Programming", "Language", "!"}
	fmt.Println("Array elements:", arr)

	s := arr[2:5]
	fmt.Println("Slice elements:", s)
	a := arr[1:4]
	fmt.Println("Another slice elements:", a)
}