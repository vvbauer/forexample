package main

import "fmt"

type square struct {
	l float64
	w float64
}

type circle struct {
	r float64
}

func (s square) area() float64 {
	a := s.l * s.w
	return a
}

func (c circle) area() float64 {
	pi := 3.14
	a := c.r * c.r * pi
	return a
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	sq1 := square{
		8.0,
		17.0,
	}

	ci1 := circle{
		20.0,
	}

	info(sq1)
	info(ci1)
}
