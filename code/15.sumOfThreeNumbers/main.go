package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1}
	res := threeSum(nums)
	fmt.Printf("%v\n", res)
}

// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 输入：nums = []
// 输出：[]
// 输入：nums = [0]
// 输出：[]
// 给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
func threeSum(nums []int) [][]int {
	//排序
	sort.Ints(nums)
	var ret [][]int
	for point := 0; point < len(nums)-2; point++ {
		//排序后首位大于零则没有满足要求的三元组
		if point == 0 && nums[point] > 0 {
			break
		}
		//每次最左边数去重
		if point > 0 && nums[point] == nums[point-1] {
			continue
		}
		left, right := point+1, len(nums)-1
		for left < right {
			sum := nums[point] + nums[left] + nums[right]
			if sum == 0 {
				ret = append(ret, []int{nums[point], nums[left], nums[right]})
				//针对左右指针得数进行去重，保证三元组每个组合唯一
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ret
}

////二维数组去重，通过map去重
//m := make(map[string]struct{})
//for i := 0; i < len(ret); i++ {
//	marshal, _ := json.Marshal(ret[i])
//	fmt.Println(string(marshal))
//	// go 序列化后是 byte[]，需要转成 string
//	m[string(marshal)] = struct{}{}
//}
//ans := [][]int{}
//for key, _ := range m {
//	var item []int
//	json.Unmarshal([]byte(key), &item)
//	ans = append(ans, item)
//}
