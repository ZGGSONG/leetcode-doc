package main

import (
	"fmt"
)

func main() {
	//digits := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	digits := []int{9}
	res := plusOne(digits)
	fmt.Println("return: ", res)
}

//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 { //满足要求切未满10则直接返回
			return digits
		}
		digits[i] = 0
	}
	digits[0] = 1 //说明全为9，需要首位改1，末位补零
	digits = append(digits, 0)
	return digits
}
