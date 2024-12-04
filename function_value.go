package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}
func main() {

	var operation func(int, int) int // tạo biến operation có kiểu func
	operation = sum                  // gán func sum cho operation

	fmt.Println(operation(5, 5))

}
