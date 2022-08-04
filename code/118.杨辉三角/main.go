package main

import (
	"fmt"
)

func main() {
	numRows := 5
	res := generate(numRows)
	fmt.Printf("Num rows: %d\n", res)
}

func generate(numRows int) [][]int {
	ret := [][]int{}
	for i := 0; i < numRows; i++ {
		row := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || i == j {
				row = append(row, 1)
				continue
			}
			row = append(row, ret[i-1][j-1]+ret[i-1][j])
		}
		ret = append(ret, row)
	}
	return ret
}

//输入: numRows = 5
//输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//输入: numRows = 1
//输出: [[1]]
func generate1(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if i == j || j < 1 {
				ret[i][j] = 1
				continue
			}
			ret[i][j] = ret[i-1][j] + ret[i-1][j-1]
		}
	}
	return ret
}
