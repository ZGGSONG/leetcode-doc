package main

import "fmt"

func main() {
	nums := []int{4, 4, 7, 6, 7}
	res := minSubsequence(nums)
	fmt.Println(res)
}

//输入：nums = [4,3,10,9,8]
//输出：[10,9]
//解释：子序列 [10,9] 和 [10,8] 是最小的、满足元素之和大于其他各元素之和的子序列。但是 [10,9] 的元素之和最大
//输入：nums = [4,4,7,6,7]
//输出：[7,7,6]
//解释：子序列 [7,7] 的和为 14 ，不严格大于剩下的其他元素之和（14 = 4 + 4 + 6）。因此，[7,6,7] 是满足题意的最小子序列。注意，元素按非递增顺序返回
//输入：nums = [6]
//输出：[6]
func minSubsequence(nums []int) []int {
	//冒号排序
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	sum1, sum2, sep := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			sum1 += nums[j]
		}
		for k := sep + 1; k < len(nums); k++ {
			sum2 += nums[k]
		}
		sep++
		if sum1 > sum2 {
			return nums[:sep]
		}
		sum1, sum2 = 0, 0
	}
	return nil
}
