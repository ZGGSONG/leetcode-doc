package main

import "fmt"

func main() {
	nums1 := []int{1, 0}
	nums2 := []int{0}
	m, n := 1, 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

// 思路：[进阶] o(m+n) 倒着遍历nums1，各自取大数放nums1中
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n; m > 0 && n > 0; i-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[i-1] = nums2[n-1]
			n--
		} else {
			nums1[i-1] = nums1[m-1]
			m--
		}
	}
	//如果只读完了nums1(大数)，需要后续补上nums2(小数)
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]
// 解释：需要合并 [1,2,3] 和 [2,5,6] 。
// 合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素
// 思路：[基础] o(n^2)合并数组->排序
func merge1(nums1 []int, m int, nums2 []int, n int) {
	x := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[x]
		x++
	}
	for i := 0; i < m+n-1; i++ {
		for j := 0; j < m+n-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				t := nums1[j]
				nums1[j] = nums1[j+1]
				nums1[j+1] = t
			}
		}
	}
}
