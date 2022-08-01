package main

import "fmt"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

*/
func main() {
	var nums []int = []int{2, 7, 11, 15}
	var target = 9
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := tmp[another]; ok {
			return []int{tmp[another], i}
		}
		tmp[nums[i]] = i
	}
	return nil
}

//xx
func twoSumOld(nums []int, target int) []int {
	var res []int
	for k, v := range nums {
		for i, j := range nums {
			if v+j == target && i != k {
				res = append(res, i)
				res = append(res, k)
			}
		}
	}
	return res
}
