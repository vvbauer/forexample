package main

import "fmt"

func main() {
	v1 := foo()
	v2, v3 := bar()
	fmt.Println(v1, v2, v3)
}

func foo() int {
	x := 42
	return x
}

func bar() (int, string) {
	return 1986, "Highway to Hell!"
}
