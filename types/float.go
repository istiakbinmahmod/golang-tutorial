package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

// untyped const
const LOL = 1000000000000000000000000000000000000000000000000000000000000000000

func main() {
	// roundError()
	// wholeNumbers()
	// formattingPrint()
	bigNumbers()
}

func bigNumbers() {
	var num uint = 4e3
	fmt.Println(num)
	num = 24e17
	fmt.Println(num)
	distance := big.NewInt(20000)
	fmt.Println(distance)
	distance = new(big.Int)
	distance.SetString("24000000000000000000000000000000000000000000000", 10)
	fmt.Println(distance)
	const division = LOL / 240000000000000000000000000000000000000000000000
	fmt.Printf("Type of %T and value %[1]v\n", division)
	/*
		Big packages
			big.Int
			big.Float
			big.Rat (like 1/3)
	*/
}

func formattingPrint() {
	num := 123
	fmt.Printf("%08b\n", num)
	future := time.Unix(20000000000, 0)
	fmt.Printf("Time %v and type %[1]T", future)
}

func roundError() {
	piggyBank := 0.0
	for count := 1; count <= 11; count++ {
		piggyBank += 0.1
	}
	fmt.Printf("%.2f\n", piggyBank)
	piggy := 0.1
	piggy += 0.2
	fmt.Println(math.Abs(piggy-0.3) < 0.0001)
}

func wholeNumbers() {
	var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("%x %x %x\n", red, green, blue)
	fmt.Printf("color: #%02x%02x%02x\n", red, green, blue)
	/*
		signed integer types
			int
			int8
			int16
			int32
			int64
	*/
	/*
		unsigned integer types
			uint
			uint8
			uint16
			uint32
			uint64
	*/
	/*
		Equivalent statements
			var minute = 60
			var minute int = 60
			minute := 60
	*/
}
