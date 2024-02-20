package main

import "fmt"

type Location struct {
	long float32
	lat  float32
}

func main() {
	var l = Location{12.01, 877.7}
	l.long = 9999
	fmt.Println(l)
}
