package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}

// 开辟一个新空间存已经遍历过的值，后面的再去与前面的值进行比较，是否满足和为target，否则继续存，而因为map的时间复杂度为O(1)，所以比使用数组进行查询更好O(n)
func twoSum(nums []int, target int) []int {
	ret := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if key, ok := ret[target-nums[i]]; ok {
			return []int{key, i}
		}
		ret[nums[i]] = i
	}
	return nil
}
