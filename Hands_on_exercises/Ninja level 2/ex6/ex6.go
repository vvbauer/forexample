package main

import "fmt"

const (
	_          = iota
	year_first = 2020 + iota
	year_second
	year_third
	year_fourth
)

func main() {
	fmt.Println(year_first, year_second, year_third, year_fourth)
}
