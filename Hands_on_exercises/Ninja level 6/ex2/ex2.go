package main

import "fmt"

func main() {

	var1 := []int{10, 15, 98, 1, 2, 17, 165}
	res1 := foo(var1...)
	fmt.Println("Result of foo is", res1)

	var2 := []int{10, 15, 98, 1, 2, 17, 165}
	res2 := bar(var2)
	fmt.Println("Result of bar is", res2)

}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
