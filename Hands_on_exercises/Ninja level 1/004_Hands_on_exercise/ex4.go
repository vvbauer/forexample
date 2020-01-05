package main

import "fmt"

type racoon int

var x racoon

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v", x)
}
