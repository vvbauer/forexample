package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	if s.ltk == true {
		fmt.Println("I am", s.first, s.last)
	}
}

func main() {
	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa2 := secretAgent{
		person{
			"Steve",
			"Rogers",
		},
		false,
	}
	sa1.speak()
	sa2.speak()
}
