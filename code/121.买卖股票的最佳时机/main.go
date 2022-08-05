package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 8, 9, 6, 19, 5, 17}
	res := maxProfit(prices)
	fmt.Printf("Max Profit: %d\n", res)
}

//单调栈
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	stack, res := []int{prices[0]}, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > stack[len(stack)-1] {
			stack = append(stack, prices[i])
		} else {
			index := len(stack) - 1
			for ; index >= 0; index-- {
				if stack[index] < prices[i] {
					break
				}
			}
			stack = stack[:index+1]
			stack = append(stack, prices[i])
		}
		if res < stack[len(stack)-1]-stack[0] {
			res = stack[len(stack)-1] - stack[0]
		}
	}
	return res
}

//模拟动态规划DP
func maxProfit2(prices []int) int {
	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
		//max = int(math.Max(float64(max), float64(prices[i]-min)))//更耗时
		//min = int(math.Min(float64(min), float64(prices[i])))
	}
	return max
}

//输入：[7,1,5,3,6,4]
//输出：5
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//输入：prices = [7,6,4,3,1]
//输出：0
//解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
//运行超时o(n^2)
func maxProfit1(prices []int) int {
	min := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i]-prices[j] < min {
				min = prices[i] - prices[j]
			}
		}
	}
	return -min
}
