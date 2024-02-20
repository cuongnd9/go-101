package main

import "fmt"

func sum(c chan int, list []int) {
	var total int
	for _, item := range list {
		total = total + item
	}
	c <- total
}

func main() {
	c := make(chan int, 10)
	list := []int{9, 10, 3}
	go sum(c, list)
	fmt.Println(<-c)
}
