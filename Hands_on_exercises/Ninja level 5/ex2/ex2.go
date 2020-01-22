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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.fav {
			fmt.Println(i, val)
		}
	}
}
