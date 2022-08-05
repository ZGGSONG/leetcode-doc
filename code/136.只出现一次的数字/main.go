package main

import "fmt"

func main() {
	nums := []int{2, 3, 3}
	res := singleNumber(nums)
	fmt.Println(res)
}

//输入: [2,2,1]
//输出: 1
//输入: [4,1,2,1,2]
//输出: 4
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//线性时间复杂度-o(n) 不能开辟新空间
//异或刚好满足只有不同时才为true
func singleNumber(nums []int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		ret = ret ^ nums[i]
	}
	return ret
}
