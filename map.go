package main

import (
	"fmt"
)

// Map : tập hợp cặp key-value
func main() {
	// Khai báo map
	mapp := make(map[int]string)

	mapp[1] = "A"
	mapp[2] = "B"
	mapp[3] = "C"
	mapp[4] = "D"

	// truy cập phần tử:
	fmt.Println(mapp[3])

	// duyệt phần tử bẳng range

	for key, value := range mapp {
		fmt.Printf("key:%d, value:%s\n", key, value)
	}
}
