package main

import ("fmt"
"github.com/WilsonLi33092/testCICD/add"
"github.com/WilsonLi33092/testCICD/subtract"
)

func main() {
	fmt.Println("Hey this works")
	fmt.Println(add.Add(2,3))
	fmt.Println(subtract.Subtract(3,1))
}
