package main

import "fmt"

func main() {
	defer ironMan()
	captain()
}

func ironMan() {
	fmt.Println("Good morning, Mr. Stark")
}

func captain() {
	fmt.Println("Hello Steve")
}
