package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32 = 10

	fmt.Printf("P stores %v\n", p)
	fmt.Printf("P refrences %v\n", *p)
	fmt.Printf("I stores %v\n", i)
	p = &i
	fmt.Printf("P stores %v\n", p)
	fmt.Printf("P refrences %v\n", *p)
	fmt.Printf("I stores %v\n", i)
	*p = 25
	fmt.Printf("P stores %v\n", p)
	fmt.Printf("P refrences %v\n", *p)
	fmt.Printf("I stores %v\n", i)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("thing1 memory location is %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("Result = %v\n", result)
	fmt.Printf("thing 1 =%v \n", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("thing2 memory location is %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
