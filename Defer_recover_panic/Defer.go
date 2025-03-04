package main

import (
	"fmt"
)

func ex1() {
	fmt.Println("1")
	defer fmt.Println("2") // Defer function will run after ex1 func finishes
	fmt.Println("3")
}

func ex2() {
	//Defer will be sorted a stack to return
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}

func ex3() {
	i := 10
	defer fmt.Println(i) // defer will be save value of variable
	i = 20
}

func main() {
	//ex1()
	//ex2()
	ex3()
}
