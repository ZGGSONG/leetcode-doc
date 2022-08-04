package main

import (
	"fmt"
)

func main() {
	numRows := 3
	res := getRow(numRows)
	fmt.Printf("Num rows: %d\n", res)
}

/**
index lenth column numbers
0     1     1      1
1     2     2      1,1
2     3     3      1,2,1
3     4     4      1,3,3,1
4     5     5      1,4,6,4,1
5     6     6      1,5,10,10,5,1
6     7     7      1,6,15,20,15,6,1
7,    8,    8      1,7,21,35,35,21,7,1
*/
func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	ret[0] = 1
	for i := 1; i <= rowIndex; i++ {
		ret[i] = ret[i-1] * (rowIndex - i + 1) / i
	}
	return ret
}

//输入: rowIndex = 3
//输出: [1,3,3,1]
//输入: rowIndex = 0
//输出: [1]
//输入: rowIndex = 1
//输出: [1,1]
func getRow1(rowIndex int) []int {
	ret := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		ret[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if i == j || j < 1 {
				ret[i][j] = 1
				continue
			}
			ret[i][j] = ret[i-1][j] + ret[i-1][j-1]
		}
	}
	return ret[rowIndex]
}
