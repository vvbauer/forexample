package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(len(x))
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
}
