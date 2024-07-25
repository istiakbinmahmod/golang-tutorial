package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// main1()
	// main2()
	// main3()
	main4()
}

func main1() {
	var myName = "Istiak Bin Mahmod"
	fmt.Println("myName contains Istiak: ", strings.Contains(myName, "Istiak"))
	fmt.Println("myName contains Masum: ", strings.Contains(myName, "Masum"))
	var age = 41
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	if strings.Contains(myName, "Istiak") {
		fmt.Println("LOL")
	} else {
		fmt.Println("BAL")
	}
}

func main2() {
	var year = 1900
	// short circuit logic
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		fmt.Println("The year", year, "is a leap year")
	} else {
		fmt.Println("The year", year, "is not a leap year")
	}
}

func main3() {
	var str = "Istiak Bin Mahmod"
	// switch strings.Contains(str, "Istiak") {
	// case true:
	// 	fmt.Println("str contains Istiak")
	// case false:
	// 	fmt.Println("str does not contain Istiak")
	// default:
	// 	fmt.Println("There really shouldn't any case like this")
	// }
	switch {
	case strings.Contains(str, "Istiak"):
		fmt.Println("str contains Istiak")
		fallthrough
	default:
		fmt.Println("str does not contain Istiak")
	}
}

func main4() {
	var count = 1

	for {
		if count > 10 {
			break
		}
		fmt.Println(count)
		time.Sleep(time.Second)
		count++
	}
}
