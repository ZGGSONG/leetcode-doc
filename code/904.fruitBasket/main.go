package main

import "fmt"

func main() {
	fruits := []int{1}
	res := totalFruit(fruits)
	fmt.Println(res)
}

// totalFruit
//
//	@Description:滑动窗口方法
//	@param []int
//	@return int
func totalFruit(fruits []int) int {
	totalMap := make(map[int]int) //定义一个map存储符合条件的数
	slowIndex, fastIndex := 0, 0  //定义双指针指针
	ret := 0                      //定义一个值来存储符合条件子数组的最大长度
	for ; fastIndex < len(fruits); fastIndex++ {
		totalMap[fruits[fastIndex]]++ //向map中存储数
		for len(totalMap) > 2 {       //当map中key的数量超过两个则可以计算子数组
			tmp := fastIndex - slowIndex //计算当前符合条件子数组长度
			if tmp > ret {               //如果当前长度大于之前的最大长度，则更新最大长度
				ret = tmp
			}
			totalMap[fruits[slowIndex]]--
			if totalMap[fruits[slowIndex]] == 0 { //当map fruits[slowIndex]的value减1后等于零则删除这个键
				delete(totalMap, fruits[slowIndex])
			}
			slowIndex++ //慢指针向后移动
		}
	}
	tmp := fastIndex - slowIndex //快指针读完整个数组后再次判断快慢指针之间的长度是否大于最大长度
	if tmp > ret {
		ret = tmp
	}
	return ret
}
