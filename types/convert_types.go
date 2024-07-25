package main

import (
	"fmt"
	"strconv"
)

/*
Convert between numeric, string, and Boolean types
*/
func main() {
	var ans bool
	error := false
	str := "bruh"
	switch str {
	case "true", "yes", "1":
		ans = true
	case "false", "no", "0":
		ans = false
	default:
		error = true
	}
	if error {
		fmt.Println("Error happened during converting string to boolean")
	} else {
		fmt.Println("Converted value", ans)
	}

	num := 123
	numToStr := strconv.Itoa(num)
	fmt.Printf("numToStr: %v\n", numToStr)

	numToStr = "1c3"
	num, err := strconv.Atoi(numToStr)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("num : %v\n", num)
	}

	num = 66
	// numToStr = string(num)
	numToStr = fmt.Sprintf("%v", num)
	fmt.Printf("numToStr: %v\n", numToStr)
}
