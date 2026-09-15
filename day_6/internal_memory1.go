package main

import "fmt"

const a = 10

var p = 1000

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println("The value of z is ", z)
	}
	add(5, 4)
	add(a, p)
}

func main() {
	call()

	fmt.Println("The value of a is ", a)
}

func init() {
	fmt.Println("Hello from init function")
}


/*
	2 phases
		1. compleation phases
		2. execution phases


		************* Compilation phase *************
		1. The compiler reads the source code and checks for syntax errors.
		2. The compiler generates an intermediate representation of the code.
		3. The compiler generates machine code from the intermediate representation.
		4. The compiler creates a binary file from the machine code.
	
		************* Execution phase *************
		1. The operating system loads the binary file into memory.
		2. The operating system starts executing the binary file.	

		********************** code segment**********************

		constants and functions are allocated in the code segment(read-only).
		And all the global variables are allocated in the data segment.

		a = 10 is a constant and it is allocated in the code segment(read-only).
		call() is a function and it is allocated in the code segment(read-only).
		add() is a function and it is allocated in the code segment(read-only).
		main() is a function and it is allocated in the code segment(read-only).
		init() is a function and it is allocated in the code segment(read-only).
		
		************* data segment *************
		All the global variables are allocated in the data segment.
		p = 1000 is a global variable and it is allocated in the data segment.

		******************** stack segment ********************


		***********************heap segment**********************

		################### execution phase ####################
		The stack is used for function execution and local variables.
		When a function is called, a new stack frame is created for that function.
		The stack frame contains the local variables and the return address of the function.
		When the function returns, the stack frame is destroyed and the control is returned to the caller function.








*/







/*

	* 2 phases for the Go Program running:
	1. compilation phase
	2. execution phase

	--> go run main.go => compile--> create a binary file named main--> then automatically execute the  binary file

	--> go build main.go => compile--> create a binary file named main --> then we can execute the binary file by running ./main

	--> Go is a compiled language, so it compiles the code before executing it.
	-->In the compilation phase, all the constants and functions are allocated in the code segment(read-only). And all the global variables are allocated in the data segment.
	--> In the execution phase, the stack is used for function execution and local variables


*/