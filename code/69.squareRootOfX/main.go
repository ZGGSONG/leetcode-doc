package main

import "fmt"

func main() {
	x := 10
	res := mySqrt(x)
	fmt.Printf("%v\n", res)
}

// 解法二 牛顿迭代法
func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

// 解法一 二分, 找到最后一个满足 n^2 <= x 的整数n
func mySqrt2(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r + 1) / 2
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

func mySqrt1(x int) int {
	if x == 1 {
		return 1
	}
	left, right := 0, x
	for right-left > 1 {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid
		} else if mid*mid < x {
			left = mid
		} else {
			return mid
		}
	}
	return left
}
