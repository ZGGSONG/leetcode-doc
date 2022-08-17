package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 12
	res := search(nums, target)
	fmt.Printf("Result: %v\n", res)
}

//左闭右开
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//左闭右闭
func search1(nums []int, target int) int {
	begin, end, mid := 0, len(nums)-1, (len(nums)-1)/2
	for begin <= end {
		if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			begin = mid + 1
		} else {
			return mid
		}
		mid = (begin + end) / 2
	}
	return -1
}
