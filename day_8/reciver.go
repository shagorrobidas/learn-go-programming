package main

import "fmt"

type Login struct {
	Username string
	Password string
}

func checkLogin(login Login) bool {
	if login.Username == "shagor" && login.Password == "123456" {
		return true
	}
	return false
}

func (login Login) CheckLoginDetails() bool {
	return checkLogin(login)
}

func main() {
	var login Login
	
	fmt.Print("Enter your username: ")
	fmt.Scanln(&login.Username)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&login.Password)

	if login.CheckLoginDetails() {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Invalid username or password.")
	}
}