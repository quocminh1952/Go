package main

// Tại sao phải dùng interface
// => xây dựng các hàm có logic giống nhau chỉ khác nhau về kiểu dữ liệu truyền vào

import (
	"fmt"
)

type englishBot struct{}
type vietBot struct{}

// tạo 1 interface có method abstract là Talk
// các struct nào implement tất cả method của interface thì struct đó đã implement interface
type Bot interface {
	Talk() string
}

// Định nghĩa method Talk() cho englishBot
func (enBot englishBot) Talk() string {
	return "Hello , wellcome to England"
}

// Điịnh nghĩa method Talk() cho VietBot
func (vnBot vietBot) Talk() string {
	return "Xin chao, chao mung toi Viet Nam"
}

// printlnTalk() chỉ cho phép các implement của Bot sử dụng
func printlnTalk(bot Bot) {
	fmt.Println(bot.Talk())
}

func main() {
	// tạo 1 instance của vietBot
	vnBot := vietBot{}
	// tạo 1 instance của englishBot
	enBot := englishBot{}

	// Vì vnBot và enBot đều cung cấp tất cả các phương thức có trong interface Bot
	// Nên vnBot và enBot đều là 1 implement của Bot
	printlnTalk(enBot)
	printlnTalk(vnBot)
}
