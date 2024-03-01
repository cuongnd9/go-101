package main

import "fmt"

func main() {
	fmt.Printf("Hello %s %s \n", "Harry", "Tran")
	var a int = 10
	var b *int
	fmt.Println("b is ", b)
	b = &a
	fmt.Println("b is ", b)
	fmt.Println("b is ", *b)
	*b = 30
	fmt.Println("b is ", b)
	fmt.Println("b is ", *b)
	fmt.Println("a is ", a)
}
