package main

import "fmt"

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	duplicates := removeDuplicates(nums)
	fmt.Println(duplicates)
}

// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4 | 2,2,3,3,4]
// 说明：用后面未重复部分覆盖重复部分
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}
