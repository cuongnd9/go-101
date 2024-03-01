package main

import "fmt"

func main() {
	var i1, i2 interface{} // Khai báo hai biến interface mà không gán giá trị
	var i3 *int

	// So sánh i1 với nil
	if i1 == nil {
		fmt.Println("i1 là nil")
	} else {
		fmt.Println("i1 không là nil")
	}

	// So sánh i2 với nil
	if i2 == nil {
		fmt.Println("i2 là nil")
	} else {
		fmt.Println("i2 không là nil")
	}

	// So sánh i1 với i2
	if i1 == i2 {
		fmt.Println("i1 == i2")
	} else {
		fmt.Println("i1 != i2")
	}

	// So sánh i1 với i3
	if i1 == i3 {
		fmt.Println("i1 == i3")
	} else {
		fmt.Println("i1 != i3")
	}
}
