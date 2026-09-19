package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Login struct {
	Username string
	Password string
}


func main() {
	var login Login
	
	fmt.Print("Enter your username: ")
	fmt.Scanln(&login.Username)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&login.Password)

	if login.Username == "shagor" && login.Password == "123456" {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Invalid username or password.")
	}

	// var user User
	

	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&user.Name)

	// fmt.Print("Enter your age: ")
	// fmt.Scanln(&user.Age)

	// fmt.Println("\nUser Information:")
	// fmt.Println("Name:", user.Name)
	// fmt.Println("Age:", user.Age)


	// var user1 User

	// user1 = User{  // instance or object  of User struct 
	// 	Name : "Shagor",
	// 	Age : 30,
	// }
	// fmt.Println("First User information:")
	// fmt.Println("Name:", user1.Name)
	// fmt.Println("Age:", user1.Age)

	// user2 := User{
	// 	Name : "Robidas",
	// 	Age : 24,
	// }
	// fmt.Println("Second User information:")
	// fmt.Println("Name:", user2.Name)
	// fmt.Println("Age:", user2.Age)

	// The main function is empty and does not perform any operations.
}