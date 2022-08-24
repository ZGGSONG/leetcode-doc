package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 3, 3, 4, 5, 6}
	val := 3
	res := removeElement(nums, val)
	fmt.Printf("removeElement: %v\n", res)
}

func removeElement(nums []int, val int) int {
	var slowIndex int
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}
