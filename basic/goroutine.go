package main

import (
	"fmt"
	"sync"
)

func sum(c chan int, list []int) {
	var total int
	for _, item := range list {
		total = total + item
	}
	c <- total
}

func sayHello(name string, wg *sync.WaitGroup) {
	fmt.Println("Hello ", name)
	defer wg.Done()
}

func main() {
	c := make(chan int, 10)
	list1 := []int{9, 10, 3}
	list2 := []int{9, 10, 5}
	go sum(c, list1)
	go sum(c, list2)
	fmt.Println(<-c, <-c)

	var wg sync.WaitGroup
	wg.Add(3)
	go sayHello("Cuong", &wg)
	go sayHello("Harry", &wg)
	go sayHello("Thuong", &wg)
	wg.Wait()
	fmt.Println("END")
}
