package main
import "fmt"

// var name string
// var age int

// func welcomeName(name string){
// 	fmt.Println("Welcome ",name)
// }

// func getName(name string) string{
// 	fmt.Print("Enter Your Name : ")
// 	fmt.Scan(&name)
// 	return name
// }

// func getAge(age int) int {
// 	fmt.Print("Enter Your Age: ")
// 	fmt.Scan(&age)
// 	return age
// }

// func displayName(name string){
// 	fmt.Println("Name: ", name)
// }

// func displayAge(age int){
// 	fmt.Println("You are  ", age, "Years Old")
// }

var(
	a = 20
	b = 30
)

func add(x int, y int){
	z := x + y
	fmt.Println("first value : ", x , "second value : ", y ,"Sum : ", z)
}

func main(){
	// name = getName(name)
	// age = getAge(age)
	// welcomeName(name)
	// displayAge(age)

	p := 30
	q := 40
	add(p,q) // 70
	add(a,b) // 50
	add(a,p) // 50
	add(b,q) // 70
}
