package main

import (
	"fmt"
)

func main() {
	s := "bbaa"
	t := "aba"
	res := minWindow(s, t)
	fmt.Printf("子串: %v\n", res)
}

//func minWindow(s string, t string) string {
//	ori, cnt := map[byte]int{}, map[byte]int{}
//	for i := 0; i < len(t); i++ {
//		ori[t[i]]++
//	}
//
//	sLen := len(s)
//	len := math.MaxInt32
//	ansL, ansR := -1, -1
//
//	check := func() bool {
//		for k, v := range ori {
//			if cnt[k] < v {
//				return false
//			}
//		}
//		return true
//	}
//	for l, r := 0, 0; r < sLen; r++ {
//		if r < sLen && ori[s[r]] > 0 {
//			cnt[s[r]]++
//		}
//		for check() && l <= r {
//			if r-l+1 < len {
//				len = r - l + 1
//				ansL, ansR = l, l+len
//			}
//			if _, ok := ori[s[l]]; ok {
//				cnt[s[l]] -= 1
//			}
//			l++
//		}
//	}
//	if ansL == -1 {
//		return ""
//	}
//	return s[ansL:ansR]
//}

//
// minWindow
//  @Description: 滑动窗口
//  @param s
//  @param t
//  @return string
//
func minWindow(s string, t string) string {
	var sFreq, tFreq, ret = map[byte]int{}, map[byte]int{}, "" //定义哈希来存储s,t中各字符数量以及默认返回值
	for _, v := range t {                                      //存储t字符以及数量
		tFreq[byte(v)]++
	}
	check := func(sFreq, tFreq map[byte]int) bool {
		for k, v := range tFreq {
			if sFreq[k] < v {
				return false
			}
		}
		return true
	}
	for left, right := 0, 0; right < len(s); right++ {
		sFreq[s[right]]++          //哈希存储s子串各字符以及数量
		if right-left+1 < len(t) { //长度比t还小则忽略的子串
			continue
		}
		for check(sFreq, tFreq) { //判断子串个字符数量是否不小于t各字符数量
			if right-left+1 < len(ret) || ret == "" {
				ret = s[left : right+1] //当子串长度小于当前最小子串或者当前最小子串为默认值时，更新最小子串
			}
			sFreq[s[left]]-- //子串最左侧字符数量减一
			left++           //子串左指针右移
		}
	}
	return ret
}

//
//func check(sFreq, tFreq map[byte]int) bool {
//	for k, v := range tFreq {
//		if sFreq[k] < v {
//			return false
//		}
//	}
//	return true
//}
