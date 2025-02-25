package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run() error {

	// Initilizacion of the calculator reader
	reader := bufio.NewReader(os.Stdin)

	// FIRST VALUE
	fmt.Println("Type first value")

	input1, _ := reader.ReadString('\n')

	num1, err := strconv.ParseFloat(input1[:len(input1)-1], 64)

	if err != nil {
		fmt.Printf("Please enter a valid number: %v", err)
	}
	fmt.Println("Value received: ", num1)

	// OPERATOR
	fmt.Println("Type operator (-, +, /, *)")

	input2, _ := reader.ReadString('\n')

	operator := input2[:len(input2)-1]
	// if !(operator == '+' || operator == '-' || operator == '/' || operator == '*') {
	// 	fmt.Println("Please enter a valid operator")
	// }

	fmt.Println("Value received: ", operator)

	// THIRD VALUE
	fmt.Println("Type second value")

	input3, _ := reader.ReadString('\n')

	num2, err := strconv.ParseFloat(input3[:len(input3)-1], 64)

	if err != nil {
		fmt.Printf("Please enter a valid number: %v", err)
	}
	fmt.Println("Value received: ", num2)

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		if num2 == 0 {
			return fmt.Errorf("division by zero")
		}
		result = num1 / num2
	case "*":
		result = num1 * num2
	default:
		fmt.Println("Invalid operator: ", operator)
	}

	fmt.Println("Result: ", result)
	return nil
}
