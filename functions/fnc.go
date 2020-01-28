package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Flemming")
	fmt.Println(x, y)
}

//basic func
func foo() {
	fmt.Println("hello from foo")
}

//arguments
func bar(s string) {
	fmt.Println("Hello", s)
}

//return
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

//multiple return
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false
	return a, b
}
