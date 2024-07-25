package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Learning about constants and variables in Go!")

	const PI = 3
	var (
		radius = 2
		area   = 0
	)
	area = PI * radius * radius
	fmt.Printf("area: %v\n", area)

	area++
	fmt.Printf("area: %v\n", area)

	// ++area
	// throws error for prefix increment

	var randomNumber = rand.Intn(10) + 1 // generates pseudorandom number between 1 and 10
	fmt.Printf("randomNumber: %v\n", randomNumber)

	const DISTANCE_TO_MARS = 56000000
	const DAYS_TO_REACH = 28
	var hourToReach = DAYS_TO_REACH * 24
	var speed = DISTANCE_TO_MARS / hourToReach
	fmt.Printf("Speed of ship to reach mars in 28 days is %v km/h\n", speed)
}
