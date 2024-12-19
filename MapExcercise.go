package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	// tách chuỗi thành một slice chứa các từ dựa trên khoảng trắng
	words := strings.Fields(s)

	// tạo 1 map để lưu số lần xuất hiện mỗi từ
	Count := make(map[string]int)

	// duyệt qua slice words và thêm vào map
	// khi truy cập hoặc tăng giá trị value của một key không có trong map
	// => map tự động thêm key đó vào trong nó
	for _, word := range words {
		Count[word]++ // tăng value của key : word
	}
	return Count
}

func main() {
	S := "Welcome to Go Word Excercise! Go"
	fmt.Println(WordCount(S))

}
