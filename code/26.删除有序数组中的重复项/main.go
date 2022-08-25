package main

import "fmt"

func main() {
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := []int{0}
	duplicates := removeDuplicates(nums)
	fmt.Println(duplicates)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	var slowIndex int
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[slowIndex] != nums[fastIndex] {
			slowIndex++
			nums[slowIndex] = nums[fastIndex]
		}
	}
	return slowIndex + 1
}

//                           1
//0, 1, 2, 3, 4, 2, 2, 3, 3, 4
//            1
