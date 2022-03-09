package main

import "fmt"

func main() {
	//variable declaration
	var operator string
	var number1, number2 int

	//User inputs
	fmt.Print("Enter first number: ")
	fmt.Scanln(&number1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&number2)

	fmt.Print("Please enter Operator (+,-,/,%,*):")
	fmt.Scanln(&operator)

	//default value of output
	output := 0

	//function acts based on operator selected
	switch operator {

	//addition
	case "+":
		output = number1 + number2
		break

	//subtraction
	case "-":
		output = number1 - number2
		break

	//multiplication
	case "*":
		output = number1 * number2
		break

	//division
	case "/":
		output = number1 / number2
		break

	//remainder
	case "%":
		output = number1 % number2
		break

	//default if missing entry
	default:
		fmt.Println("Invalid Operation")
	}

	//print result
	fmt.Printf("%d %s %d = %d", number1, operator, number2, output)
}
