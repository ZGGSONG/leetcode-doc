package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 1}
	target := 4
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}

//O(nlog(n))
func minSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}

// O(n)
func minSubArrayLen2(target int, nums []int) int {
	var i, j, sum, ret, flag = 0, 0, 0, math.MaxInt, true
	//遍历数组
	for j < len(nums) {
		if flag {
			sum += nums[j]
		}
		if sum >= target {
			tmp := j - i + 1
			if tmp < ret {
				ret = tmp
			}
			sum -= nums[i]
			i++
			flag = false
		} else {
			j++
			flag = true
		}
	}
	//如果没有符合条件
	if ret == math.MaxInt {
		ret = 0
	}
	return ret
}

//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//滑动窗口（另一种双指针）
//FAIL: 超时
func minSubArrayLen1(target int, nums []int) int {
	ret := math.MaxInt
	slowIndex, fastIndex, sum := 0, 0, 0
	for fastIndex < len(nums) {
		for i := slowIndex; i <= fastIndex; i++ {
			sum += nums[i]
		}
		if sum >= target {
			tmp := fastIndex - slowIndex + 1
			if tmp < ret {
				ret = tmp
			}
			slowIndex++
		} else {
			fastIndex++
		}
		sum = 0
	}
	//没有
	if ret == math.MaxInt {
		ret = 0
	}
	return ret
}

//func getSum(nums []int, start, end int) int {
//	var sum int
//	for i := start; i <= end; i++ {
//		sum += nums[i]
//	}
//	return sum
//}
