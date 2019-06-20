package main

import "fmt"

const (
	year1 = 2016 + iota
	year2 = 2016 + iota
	year3 = 2016 + iota
	year4 = 2016 + iota
)

func main() {
	fmt.Println(year1, year2, year3, year4)
}
