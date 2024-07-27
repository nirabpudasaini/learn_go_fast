package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint16 {
	return uint16(e.gallons) * uint16(e.mpg)
}

func (e electricEngine) milesLeft() uint16 {
	return uint16(e.mpkwh) * uint16(e.kwh)
}

type engine interface {
	milesLeft() uint16
}

func canMakeIt(e engine, miles uint16) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("Think about refuelling")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{15, 25}
	fmt.Println(myEngine)
	fmt.Printf("Miles left = %v \n", myEngine.milesLeft())
	canMakeIt(myEngine, 400)

	myElectricEngine := electricEngine{25, 25}
	canMakeIt(myElectricEngine, 400)
}
