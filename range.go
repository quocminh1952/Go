package main

import (
	"fmt"
)

// Range : là cách duyệt qua các phần tử cúa slice/array/map/string ...
// -> Khi sử dụng vơi slice / array : trả về index và value của mỗi phần tử
// -> với map : trả về key value
// -> với string : trả về index của mỗi ký tự và giá trị mã unicode của nó
func main() {

	var num = [5]int{1, 3, 4, 5, 4}
	slice := num[1:4]

	for index, value := range slice {
		fmt.Println("Index: %d, Value: %d", index, value)
	}

}
