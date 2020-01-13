package main

import "fmt"

func main() {
	x := "Chris De Burgh"
	if x == "Elvis Presley" {
		fmt.Println("King of kings")
	} else if x == "Chris De Burgh" {
		fmt.Println("Moonlight and vodka takes me away midnight in Moscow is lunchtime in LA")
	} else {
		fmt.Println("don't know who is this")
	}
}
