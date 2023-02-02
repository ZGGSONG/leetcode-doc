package main

import "fmt"

func main() {
	nums := []int{1}
	res := sortedSquares(nums)
	fmt.Printf("res: %v\n", res)
}

// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]
// 排序后，数组变为 [0,1,9,16,100]
func sortedSquares(nums []int) []int {
	var length = len(nums)
	var ret = make([]int, length)
	slowIndex, fastIndex, index := 0, length-1, length-1
	for slowIndex <= fastIndex {
		slowSquare, fastSquare := nums[slowIndex]*nums[slowIndex], nums[fastIndex]*nums[fastIndex]
		if slowSquare > fastSquare {
			ret[index] = slowSquare
			slowIndex++
		} else {
			ret[index] = fastSquare
			fastIndex--

		}
		index--
	}
	return ret
}

//       1
//0,1,16,9,10
//     1

// 暴力解法
func sortedSquares1(nums []int) []int {
	for i, _ := range nums {
		nums[i] *= nums[i]
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
