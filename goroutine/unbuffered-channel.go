package main

import (
	"fmt"
)

func main() {
	// Tạo một kênh không có bộ đệm
	ch := make(chan int)

	// Goroutine gửi dữ liệu vào kênh
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // Gửi giá trị vào kênh
		}
		close(ch) // Đóng kênh sau khi gửi xong
	}()

	// Goroutine đọc dữ liệu từ kênh
	for val := range ch {
		fmt.Println("Received:", val)
	}
}
