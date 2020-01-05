package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	foo()
	fmt.Println("something more")
}

func foo() {
	fmt.Println("I'm in foo")
}
