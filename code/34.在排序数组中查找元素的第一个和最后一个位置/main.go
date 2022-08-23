package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 2, 4, 5, 6}
	target := 2
	res := searchRange(nums, target)
	fmt.Println("res:", res)
}

func searchRange(nums []int, target int) []int {
	//1. 在数组外
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}
	//2. 在数组内
	left, right := 0, len(nums)-1
	var ret []int
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			//2.1 在数组范围内且存在
			mid2 := mid
			for mid-1 >= 0 && nums[mid-1] == target {
				mid--
			}
			ret = append(ret, mid)
			for mid2+1 <= len(nums)-1 && nums[mid2+1] == target {
				mid2++
			}
			ret = append(ret, mid2)
			return ret
		}
	}
	//2.2 在数组范围内且不存在
	return []int{-1, -1}
}

func searchRange1(nums []int, target int) []int {
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}

}

// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) { // 找到第一个与 target 相等的元素
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后一个与 target 相等的元素
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}
