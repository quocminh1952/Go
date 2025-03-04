package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

func (std Student) ValueReceivers() {
	std.name = "Alice"
	std.age = 30
}

func (std *Student) PointerReceivers() {
	std.name = "May"
	std.age = 18
}

func main() {

	Jane := Student{"Jane", 20}
	// tạo một bản sao của Jane để thực hiện phương thức
	//=> không thay đổi giá trị Jane
	Jane.ValueReceivers()
	fmt.Println(Jane.age)
	// phương thức nhận địa chỉ của Jane để thực hiện thay vì dùng bản sao
	// thay đổi giá trị Jane
	Jane.PointerReceivers()
	fmt.Println(Jane.age)
}
