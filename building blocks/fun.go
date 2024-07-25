package main

import "fmt"

func getSum(nums ...int) (sum int) {
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	sum := getSum(1, 2, 3, 4, 5)
	fmt.Println(sum)
}
