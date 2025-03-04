package main

import (
	"fmt"
	"time"
)

func PrintHello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go PrintHello("Sheep")
	PrintHello("fish")
	// khi chạy main() thì chương trình đã tạo 1 goroutine cho hàm main()
	// main() goroutine kết thúc thì sẽ dừng chương trình cho dừ các goroutine khác chưa chaạy xong

	// ***NẾU khai báo thêm goroutine cho PringtHello("fish")
	// bước 1 : khởi tạo goroutine cho main()
	// bước 2 : hàm main() khởi tạo goroutine cho sheep và fish
	//			=> hàm main() hoàn thành nhiệm vụ => kết thúc chương trình
	//			=> goroutine của sheep và fish chưa kịp chạy
	//			=> chương trình không in ra kết quả
}
