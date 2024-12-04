package main

import (
	"fmt"
)

func main() {

	var X int = 10
	// khai báo con trỏ
	var Y *int // giá trị hiện tại của *y = nil
	// lấy địa chỉ biến X gán cho con trỏ
	Y = &X
	// thay đổi giá trị cuả X thông qua Y
	*Y = 20
	fmt.Println(X)

}
