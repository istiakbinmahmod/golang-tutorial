package main

import "fmt"

func twoSums(nums []int, target int) (int, int) {
	lengthOfNums := len(nums)
	i, j := 0, lengthOfNums-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			return i, j
		} else if sum > target {
			j -= 1
		} else {
			i += 1
		}
	}
	return -1, -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	target := 7
	i, j := twoSums(nums, target)
	if i == -1 && j == -1 {
		fmt.Println("No two sum found")
	} else {
		fmt.Printf("The two values are (index:value) %2v:%2v and %2v:%2v\n", i, nums[i], j, nums[j])
	}
}
