package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{833, 736, 953, -584, -448, 207, 128, -445, 126, 248, 871, 860, 333, -899, 463, 488, -50, -331, 903, 575, 265, 162, -733, 648, 678, 549, 579, -172, -897, 562, -503, -508, 858, 259, -347, -162, -505, -694, 300, -40, -147, 383, -221, -28, -699, 36, -229, 960, 317, -585, 879, 406, 2, 409, -393, -934, 67, 71, -312, 787, 161, 514, 865, 60, 555, 843, -725, -966, -352, 862, 821, 803, -835, -635, 476, -704, -78, 393, 212, 767, -833, 543, 923, -993, 274, -839, 389, 447, 741, 999, -87, 599, -349, -515, -553, -14, -421, -294, -204, -713, 497, 168, 337, -345, -948, 145, 625, 901, 34, -306, -546, -536, 332, -467, -729, 229, -170, -915, 407, 450, 159, -385, 163, -420, 58, 869, 308, -494, 367, -33, 205, -823, -869, 478, -238, -375, 352, 113, -741, -970, -990, 802, -173, -977, 464, -801, -408, -77, 694, -58, -796, -599, -918, 643, -651, -555, 864, -274, 534, 211, -910, 815, -102, 24, -461, -146}
	target := -7111
	res := threeSumClosest(nums, target)
	fmt.Println(res)
}

// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// 输入：nums = [0,0,0], target = 1
// 输出：0
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closerSum := nums[0] + nums[1] + nums[2]
	for point := 0; point < len(nums)-2; point++ {
		left, right := point+1, len(nums)-1
		for left < right {
			threeSum := nums[point] + nums[left] + nums[right]
			if intAbs(threeSum-target) < intAbs(closerSum-target) {
				closerSum = threeSum
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
			}
			if threeSum > target {
				right--
			} else if threeSum < target {
				left++
			} else {
				return target
			}
		}
	}
	return closerSum
}

// intAbs
// @Description: 取int绝对值
// @param num	参数
// @return int 取绝对值目标值
func intAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
