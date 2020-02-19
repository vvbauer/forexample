package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	foo()

	func() {
		fmt.Println("Anonymous function ran")
	}()

	func(x int) {
		fmt.Println("lol", x)
	}(42)
}

func foo() {
	fmt.Println("foo ran")
}
