package main

import "fmt"

// local scope

func localScope() {
	local := 10
	fmt.Println("local", local)
}

// block
/*
	1. block -> {} curly brace
*/

func blockScope() {
	for i := 0; i < 10; i++ {
		fmt.Println("block", i)
	}
}

func main() {
	x := 18

	if x >= 18 {
		p := 10
		fmt.Println("I am matuted boy")
		fmt.Println("p value : ", p)
	}

	fmt.Println("local scope")
	localScope()

	fmt.Println("block scope")
	blockScope()

}
