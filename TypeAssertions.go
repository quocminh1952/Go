package main

import (
	"fmt"
)

// Type Asertion : Cú pháp truy cập giá trị cụ thể bên trong một interface
// và xác định kiểu của giá trị đó.

type I interface{}

func main() {
	type I interface{}
	// tạo biến interface và gán giá trị int 20
	var i I = 20
	// sử dung type asertion để kiểm tra
	// dạng không an toàn : gây panic nếu không thuộc giá trị trong ()
	// v := i.(int)
	// dạng an toàn
	v, ok := i.(string)
	fmt.Println(v, ok)
}
