package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		for y := 0; y < 3; y++ {
			fmt.Printf("%#U\n", i)
		}
	}
}

//65 to 90
