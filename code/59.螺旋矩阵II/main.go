package main

import "fmt"

func main() {
	n := 4
	res := generateMatrix(n)
	fmt.Println(res)
}

//
// generateMatrix
//  @Description: 螺旋矩阵，想象其为一个坐标系
//  @param n
//  @return [][]int
//
func generateMatrix(n int) [][]int {
	//定义开始位置，每条边的始末位置，左闭右开的偏移量，填充count
	var startX, startY, offset, count, i, j = 0, 0, 1, 1, 0, 0
	var matrix = make([][]int, n)
	for i = 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for k := 0; k < n/2; k++ {
		//上
		for j = startY; j < n-offset; j++ {
			matrix[startX][j] = count
			count++
		}
		//右
		for i = startX; i < n-offset; i++ {
			matrix[i][j] = count
			count++
		}
		//下
		for ; j > startY; j-- {
			matrix[i][j] = count
			count++
		}
		//左
		for ; i > startX; i-- {
			matrix[i][j] = count
			count++
		}
		//每次循环过后，向内层进行遍历
		startX++
		startY++
		offset++
	}
	//如果n为奇数，中间则最后在填充一位
	if n%2 == 1 {
		matrix[n/2][n/2] = count
	}
	return matrix
}
