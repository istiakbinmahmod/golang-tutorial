package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// main1()
	main2()
}

func main1() {
	for count := 0; count < 10; count++ {
		fmt.Println(count)
		time.Sleep(time.Second)
	}
}

func main2() {
	// short declaration
	if num := rand.Intn(10); num > 5 {
		fmt.Println("num is", num, "and it is greater than 5")
	} else if num2 := rand.Intn(20); num2 > 10 {
		fmt.Println("another random int:", num2)
	} else {
		fmt.Println("couldn't find such a random number")
	}
}
