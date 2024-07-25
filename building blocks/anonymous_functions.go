package main

import "fmt"

var num int = 5

func main() {
	customPrint(num)
	customPrint(num)
	anonymousFunctions()
}

func customPrint(num int) {
	fmt.Println(num)
	num++
}

func anonymousFunctions() {
	/*
		An anonymous function, also called function literal in Go, a function without any name.
	*/
	ronaldo := func(celebration string) string {
		return celebration + "ti tumi ar kedo na"
	}
	fmt.Println(ronaldo("siuu")) // you can guess the output by now :)
	func(n1, n2 int) {
		fmt.Println("Two numbers passed are:", n1, n2)
	}(10, 20) // invokes the function
}
