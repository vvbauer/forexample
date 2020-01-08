package main

import "fmt"

func main() {
	g := (42 == 42)
	h := (42 <= 42)
	i := (42 >= 42)
	j := (42 != 42)
	k := (42 < 42)
	l := (42 > 42)
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)
	fmt.Printf("%v\n", k)
	fmt.Printf("%v\n", l)
}
