package main

import "fmt"

func main() {
	var mockArr = [10]int{0, 1, 8, 19, 39, 99}
	var slice []int = mockArr[2:7]
	fmt.Println(slice)
}
