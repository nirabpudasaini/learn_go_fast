package main

import "fmt"

func main() {
	// var intArr [3]int32
	intArr := [...]int32{1, 2, 3}
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{2, 6, 8}
	fmt.Println(intSlice)
	fmt.Printf("Length: %v and Capacity: %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 10)
	fmt.Printf("Length: %v and Capacity: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{3, 5}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	intSlice3 := make([]int32, 4, 10) // create slice with length 4 and capacity 10
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Ivan": 10, "Katya": 11, "Johana": 10, "Kristen": 12}
	fmt.Println(myMap2["Ivan"])
	fmt.Println(myMap2["Jason"]) //returns default value for uint8
	var intValue, isInMap = myMap2["Jason"]
	switch isInMap {
	case true:
		fmt.Println(intValue)
	default:
		fmt.Println("Can not find the key")
	}
	delete(myMap2, "Kristen")

	for name, value := range myMap2 {
		fmt.Printf("Name: %v , Value: %v \n", name, value)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	var i int = 0
	for i < 10 {
		fmt.Print(i)
		i = i + 1
	}
	fmt.Println()

	i = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Print(i)
		i += 1
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

}
