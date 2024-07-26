package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var intNum int = 32767
	// int can be int, int8, int32 or int64
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 10.112
	// float can be float32 or float 64
	fmt.Println(floatNum)

	var result float64 = floatNum * float64(intNum)
	// operations need to be preformed on same type or need to be type casted
	fmt.Println(result)

	intNum1 := 5
	var intNum2 = 2
	// := can be used as shorthand for var keyword
	// type if not specified are infered at first assignment
	println(intNum1 / intNum2)
	// integer division results in integer
	println(intNum1 % intNum2)
	// % can be used to find quotient

	var myString string = "Hello World"
	var multiLineString = `hello 
	world
	of 
	go`
	fmt.Println(myString)
	fmt.Println(myString + " " + multiLineString)
	// + operator concanetates the strings
	fmt.Println(utf8.RuneCountInString(myString))
	// built in len function will give larger values for unicode

	var myBoolean bool = true
	fmt.Println(myBoolean)

	var1, var2 := 1, 2
	// multiple vaieables of same type can be initialized in single line
	fmt.Println(var1, var2)

	const pi float32 = 3.14
	// const are read only variables
	fmt.Println(pi)

}
