package main

import "fmt"

func main() {
	t := struct {
		brand  string
		doors  int
		color  string
		wheels int
	}{
		brand:  "Peterbilt",
		doors:  2,
		color:  "White",
		wheels: 6,
	}

	fmt.Println(t.brand)
	fmt.Println(t.doors)
	fmt.Println(t.color)
	fmt.Println(t.wheels)

}
