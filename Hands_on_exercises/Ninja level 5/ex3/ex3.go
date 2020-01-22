package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	peterbilt := truck{
		vehicle: vehicle{
			doors: 2,
			color: "yellow",
		},
		fourWheel: true,
	}

	continental := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(peterbilt)
	fmt.Println(continental)

	fmt.Println(peterbilt.color)
	fmt.Println(continental.color)
}
