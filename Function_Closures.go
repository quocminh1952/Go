package main

import (
	"fmt"
)

// hàm closures trả về một hàm ẩn danh
func Fibonacci() func() int {
	before1 := 0
	before2 := 1
	// từ lần gọi thứ 2 hàm clores chỉ chạy đoạn code trong hàm ẩn danh
	return func() int {
		current := before1
		before1 = before2
		before2 = current + before1
		return current
	}
}

func main() {
	Fi := Fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(Fi())
	}
}
