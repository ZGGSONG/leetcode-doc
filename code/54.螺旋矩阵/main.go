package main

import (
	"fmt"
	"math"
)

func main() {
	//var matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	//var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var matrix = [][]int{{6, 9, 7}}
	res := spiralOrder(matrix)
	fmt.Println(res)
}

func spiralOrder(matrix [][]int) []int {
	var rows = len(matrix)
	var columns = len(matrix[0])
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var ret []int
	var startX, startY, i, j int
	var loop, mid int
	offset := 1
	if rows > columns {
		loop = columns / 2
		mid = columns / 2
	} else {
		loop = rows / 2
		mid = rows / 2
	}
	for ; loop > 0; loop-- {
		for j = startY; j < startY+columns-offset; j++ {
			ret = append(ret, matrix[startX][j])
		}

		for i = startX; i < startX+rows-offset; i++ {
			ret = append(ret, matrix[i][j])
		}

		for ; j > startY; j-- {
			ret = append(ret, matrix[i][j])
		}

		for ; i > startX; i-- {
			ret = append(ret, matrix[i][j])
		}
		startX++
		startY++
		offset += 2
	}
	if int(math.Min(float64(rows), float64(columns)))%2 == 1 {
		if rows < columns {
			for i = mid; i < mid+columns-rows+1; i++ {
				ret = append(ret, matrix[mid][i])
			}
		} else {
			for i = mid; i < mid+rows-columns+1; i++ {
				ret = append(ret, matrix[i][mid])
			}
		}
	}
	return ret
}
