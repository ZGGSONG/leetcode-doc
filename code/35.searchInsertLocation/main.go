package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 10}
	target := 7
	location := searchInsert(nums, target)
	fmt.Printf("Location: %v\n", location)
}

// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
// 请必须使用时间复杂度为 O(log n) 的算法
func searchInsert(nums []int, target int) int {
	//长度为0或者小于所有
	if len(nums) == 0 || target < nums[0] {
		return 0
	} else if target > nums[len(nums)-1] { //大于所有
		return len(nums)
	}
	//数组中存在或者在数组大小范围内
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high + low) / 2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return low
}
