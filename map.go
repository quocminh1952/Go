package main

import (
	"fmt"
)

// Map : tập hợp cặp key-value
func main() {
	// Khởi tạo map
	mapp := make(map[int]string)
	// map_name := make(map[type_key]type_value)
	mapp2 := map[int]string{
		2: "apple",
		3: "banana",
	}
	// mapp tên maps
	// int : kiểu của key
	// string : kiểu của value
	mapp[1] = "A"
	mapp[2] = "B"
	mapp[3] = "C"
	mapp[4] = "D"

	// truy cập phần tử:
	fmt.Println(mapp[3])
	fmt.Println(mapp2)
	// duyệt phần tử bẳng range
	for key, value := range mapp {
		fmt.Printf("k ey:%d, value:%s\n", key, value)
	}
}
