package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
	target := -11
	res := fourSum(nums, target)
	fmt.Println(res)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-3; i++ {
		//去除此种状态，不能跟第三题一样，不是与0相比，多个负数相加会得到更小的值
		//if i == 0 && nums[i] > target {
		//	break
		//}
		//起始值去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			//起始值去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return ret
}
