package main

import (
	"fmt"
	"math"
)

func main() {
	// main1()
	main2()
}

func main1() {
	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	fmt.Println(pi64)
	fmt.Println(pi32)
	var price float32
	fmt.Println(price)
	fmt.Printf("formatted PI %v\n", pi64)
	fmt.Printf("formatted PI %f\n", pi64)
	fmt.Printf("formatted PI %100.4f\n", pi64)
	fmt.Printf("formatted PI %.100f\n", pi64)
}

func main2() {
	third := 1.0 / 3
	one := third + third + third
	fmt.Println(one)
	val := 0.1
	val += 0.2
	fmt.Println(val)
	farheinheit := 122.2
	celsius := (farheinheit - 32) / 9.0 * 5.0
	fmt.Println(celsius)
	celsius = (farheinheit - 32) / 9.0 * 5.0
	fmt.Println(celsius)
}
