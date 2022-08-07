package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(height)
	fmt.Println("res: ", res)
}

//时间复杂度O(n)--[双指针](https://zhuanlan.zhihu.com/p/71643340)
func maxArea(height []int) int {
	ret, tmp, left, right := 0, 0, 0, len(height)-1
	for left < right {
		if height[left] > height[right] {
			tmp = (right - left) * height[right]
			right--
		} else {
			tmp = (right - left) * height[left]
			left++
		}
		if ret < tmp {
			ret = tmp
		}
	}
	return ret
}

//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
//输入：height = [1,1]
//输出：1
//n == height.length
//2 <= n <= 105
//0 <= height[i] <= 104
//暴力解法不通过，超出时间限制 所以需要使用时间复杂度为O(n)的方法
func maxArea1(height []int) int {
	area := 0
	h := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			if height[j] < height[i] {
				h = height[j]
			} else {
				h = height[i]
			}
			if (j-i)*h > area {
				area = (j - i) * h
			}
		}
	}
	return area
}
