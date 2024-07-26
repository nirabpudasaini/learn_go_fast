package main

import (
	"errors"
	"fmt"
)

func main() {
	printName("Golang")

	num1, num2 := 5, 5
	var result, remainder, err = intDivision(num1, num2)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("Result = %v \n", result)
	default:
		fmt.Printf("Quotient = %v and Remainder = %v \n", result, remainder)

	}

	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result = %v \n", result)
	} else {
		fmt.Printf("Quotient = %v and Remainder = %v \n", result, remainder)
	}

}

func printName(name string) {
	fmt.Println("Your name is: " + name)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	quotient := numerator / denominator
	remainder := numerator % denominator
	return quotient, remainder, err
}
