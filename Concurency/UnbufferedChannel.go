package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // Tạo channel không bộ đệm

	go func() {
		fmt.Println("Sending data...")
		ch <- "Hello from goroutine" //1 Gửi dữ liệu và CHỜ
		fmt.Println("Data sent!")    //2 Thực thi sau khi dữ liệu được nhận
	}()

	time.Sleep(1 * time.Second) // Đảm bảo goroutine gửi chạy trước
	fmt.Println("Receiving data...")
	data := <-ch //2 Nhận dữ liệu từ channel
	fmt.Println("Received:", data)

}
