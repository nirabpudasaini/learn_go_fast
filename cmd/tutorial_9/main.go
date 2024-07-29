package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_MILK_PRICE float32 = 1

func main() {
	var chickenChannel = make(chan string)
	var milkChannel = make(chan string)
	var websites = []string{"rimi.ee", "prismamarket.ee", "selver.ee", "lidl.ee"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkMilkPrices(websites[i], milkChannel)
	}
	sendMessage(chickenChannel, milkChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		fmt.Printf("Chicken Website: %s, price %v \n", website, chickenPrice)
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkMilkPrices(website string, milkChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var milkPrice = rand.Float32() * 5
		fmt.Printf("Milk Website: %s, price %v \n", website, milkPrice)
		if milkPrice <= MAX_MILK_PRICE {
			milkChannel <- website
			break
		}
	}

}

func sendMessage(chickenChannel chan string, milkChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("Chicken deal found at %s \n", website)
	case website := <-milkChannel:
		fmt.Printf("Milk deal found at %s \n", website)

	}

}
