package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "and i am", p.age, "years old.")
}

func main() {
	p1 := person{
		first: "Bob",
		last:  "Silent",
		age:   42,
	}
	p1.speak()
}
