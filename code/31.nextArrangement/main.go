package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 3, 5, 7, 6, 4}
	//nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				sort.Ints(nums[i+1:])
				return
			}
		}
	}
	sort.Ints(nums)
}
