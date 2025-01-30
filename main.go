package main

import "fmt"

func main() {
	fmt.Println("Hey this works")
	fmt.Println(add(2,3))
}

func add(x,y int) int {
	return x+y
}