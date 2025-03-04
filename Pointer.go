package main

import (
	"fmt"
)

// Pointer
func main() {

	//Khai báo biến a
	a := 23

	//khai báo pointer b
	var b *int //
	// Giá trị của con trỏ b khi chưa được gán là : nil
	fmt.Printf("giá trị b = %v\n kiểu của b:%T\n", b, b)

	// sử dụng b để trỏ tới a
	i := &a
	fmt.Printf("giá trị i = %v\n kiểu của i:%T", i, i)

}
