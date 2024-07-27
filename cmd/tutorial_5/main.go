package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "निरब" // Strings are in UTF-8
	for i, v := range myString {
		fmt.Println(i, v) //Each character takes 3 bytes which is seen in index
	}
	fmt.Printf("Length: %v \n", len(myString)) // Len also gives lenth of bytes

	var myRune = []int32("निरब") // Uses UTF-32
	for i, v := range myRune {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}
	fmt.Printf("Length: %v \n", len(myString))

	/*
		String are immutable so different operation can be innefficient
		Use functions from strings to work with strings
	*/
	var strSlice = []string{"a", "b", "c"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var finalStr = strBuilder.String()
	fmt.Println(finalStr)

}
