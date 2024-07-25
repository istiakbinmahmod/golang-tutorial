package main

import "fmt"

type arithmeticFunc func(int, int) int

func main() {
	// assignFunctionToVariable()
	passFunctionToFunction()
}

func assignFunctionToVariable() {
	dummyVariable := dummyFunction
	fmt.Println(dummyVariable())
}

func passFunctionToFunction() {
	arithmeticOperations(10, 5, add, subtract, multiply, divide)
}

func dummyFunction() string {
	return "Some dummy function"
}

func arithmeticOperations(n1, n2 int, add, sub, mul, div arithmeticFunc) {
	fmt.Println("4 types operations example:")
	fmt.Printf("Add operation for %v and %v and result: %v\n", n1, n2, add(n1, n2))
	fmt.Printf("Subtract operation for %v and %v and result: %v\n", n1, n2, sub(n1, n2))
	fmt.Printf("Multiply operation for %v and %v and result: %v\n", n1, n2, mul(n1, n2))
	fmt.Printf("Divide operation for %v and %v and result: %v\n", n1, n2, div(n1, n2))
}

func add(n1, n2 int) int {
	return n1 + n2
}

func subtract(n1, n2 int) int {
	return n1 - n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}

func divide(n1, n2 int) int {
	return n1 / n2
}
