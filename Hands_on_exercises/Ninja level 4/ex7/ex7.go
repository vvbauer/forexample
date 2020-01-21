package main

import "fmt"

func main() {
	j := []string{"James", "Bond", "Shaken, not stirred"}
	m := []string{"Miss", "Moneypenny", "Helloooo, James"}
	fmt.Println(j)
	fmt.Println(m)

	jm := [][]string{j, m}
	fmt.Println(jm)

	for i, xs := range jm {
		fmt.Println("record:", i)
		for x, val := range xs {
			fmt.Printf("\t index %v \t value: \t %v \n", x, val)
		}
	}
}
