package main

import (
	"fmt"
	"example.com/mathlib"
)

func main() {
	// go mod init example.com command custome package
	fmt.Println("Showing Custome package ")
	mathlib.Add(10, 20)
	mathlib.sum()

}
