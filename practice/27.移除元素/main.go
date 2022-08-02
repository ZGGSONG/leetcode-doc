package main

import "fmt"

func main() {
	nums := []int{1, 3, 3, 4}
	val := 3
	element := removeElement(nums, val)
	fmt.Println(element)
}

//输入：nums = [3,2,2,3], val = 3
//输出：2, nums = [2,2]
//注意：nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			//标记为当前=val的下标，也即所求数组长度
			j++
		}
	}
	return j
}
