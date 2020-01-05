package main

import "fmt"

func main() {
	//short declaration operator
	//Declare a variable and Assign a Value at the same time
	x := "Hello world!"

	fmt.Println(x)

	foo()

	x = "something more"

	fmt.Println(x)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}
