package main

import "fmt"

func main() {
	foo()
	bar("James")
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello", s)
}
