package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("i am", s.first, s.last, " - the secrent agent said")
}

func (p person) speak() {
	fmt.Println("i am", p.first, p.last, " - the person said")
}

func bar(h human) {
	fmt.Println("Hello from bar - the bar said")
}

func todd(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I love pasta, - said", h.(person).first)
	case secretAgent:
		fmt.Println("I hate pasta, - said", h.(secretAgent).first)
	}
	fmt.Println("You are", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	p1 := person{
		"Jabrail",
		"Seratov",
	}
	sa1.speak()
	p1.speak()
	bar(sa1)
	todd(sa1)
	todd(p1)
}
