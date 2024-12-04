package main

import (
	"fmt"
)

func main() {

	type Book struct {
		Title  string
		Author string
		Price  float64
	}

	Books := [3]Book{
		{"TOEIC 1", "MinhNguyen", 10.9},
		{"TOEIC 2 ", "TRANPhan", 11.3},
		{"TOEIC 4", "NGUYENTIEN", 12.5},
	}

	fmt.Println(Books)
}
