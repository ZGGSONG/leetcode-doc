package main

import "fmt"

func main() {
	nums := []int{-1, 0, 0, -4, 5, 6, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

//双指针
func moveZeroes(nums []int) {
	var slowIndex int
	for k, v := range nums {
		if v == 0 {
			slowIndex = k
			break
		}
	}
	for fastIndex := slowIndex; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 {
			nums[slowIndex], nums[fastIndex] = nums[fastIndex], nums[slowIndex]
			slowIndex++
		}
	}
}

//双指针
func moveZeroes1(nums []int) {
	var slowIndex int
	//是否找到第一个0
	var flag bool
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		//先找到第一个0
		if nums[fastIndex] == 0 && !flag {
			slowIndex = fastIndex
			flag = true
		}
		//找到之后跟后面非0得数进行对调
		if nums[fastIndex] != 0 && flag {
			nums[slowIndex], nums[fastIndex] = nums[fastIndex], nums[slowIndex]
			slowIndex++
		}
	}
}

//      1
//1,0,0,3,12
//  1
