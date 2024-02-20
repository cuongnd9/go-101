package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(add(9, 10))
	fmt.Println(swap(9, 19))
}
