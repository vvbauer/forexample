package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	fav   []string
}

func main() {

	p1 := person{
		first: "Peter",
		last:  "Quill",
		fav: []string{
			"Strawberry",
			"Vanilla",
			"Chocolate",
		},
	}

	p2 := person{
		first: "Steven",
		last:  "Strange",
		fav: []string{
			"Blackberry",
			"Green tea",
			"Raspberry",
		},
	}

	fmt.Println("First Name: \t", p1.first)
	fmt.Println("Last  Name: \t", p1.last)
	fmt.Println("Favorite ice cream favours:")
	for i, v := range p1.fav {
		fmt.Println(i, v)
	}

	fmt.Println("\n")

	fmt.Println("First Name: \t", p2.first)
	fmt.Println("Last  Name: \t", p2.last)
	fmt.Println("Favorite ice cream favours:")
	for i, v := range p2.fav {
		fmt.Println(i, v)
	}

}
