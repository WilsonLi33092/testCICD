package main

import ("fmt"
"github.com/WilsonLi33092/testCICD/add"
"github.com/WilsonLi33092/testCICD/subtract"
"github.com/WilsonLi33092/testCICD/calculations"
)

func main() {
	fmt.Println("Hey this works")
	fmt.Println(add.Add(2,3))
	fmt.Println(subtract.Subtract(3,1))
	fmt.Println(calculations.Multiply(3,1))
	fmt.Println(calculations.Divide(2,1))
	fmt.Println(calculations.Mod(3,2))
	fmt.Println(calculations.Exponent(2,1))
	var decision string
	fmt.Println("Enter a decision (add/subtract/divide/multiply/exponent/mod)")
	fmt.Scan(&decision)
	var num1 int
	var num2 int
	fmt.Println("Enter the first number you would like to use")
	fmt.Scan(&num1)
	fmt.Println("Enter the second number you would like to use")
	fmt.Scan(&num2)
	fmt.Printf("The first number you used is %d, The second number you used is %d, And the operation you chose is %v \n", num1, num2, decision)
	switch decision {
		case "add":
			fmt.Printf("Your result is %d", add.Add(num1, num2))
		case "subtract":
			fmt.Printf("Your result is %d", subtract.Subtract(num1,num2))
		case "divide":
			fmt.Printf("Your result is %d", calculations.Divide(num1,num2))
		case "multiply":
			fmt.Printf("Your result is %d", calculations.Multiply(num1,num2))
		case "exponent":
			fmt.Printf("Your result is %d", calculations.Exponent(num1,num2))
		case "mod":
			fmt.Printf("Your result is %d", calculations.Mod(num1,num2))
		default:
			fmt.Println("You did not select a valid option")
	}

}
