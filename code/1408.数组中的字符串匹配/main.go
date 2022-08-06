package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"leetcoder", "leetcode", "od", "hamlet", "am"}
	res := stringMatching(words)
	fmt.Println(res)
}

//输入：words = ["mass","as","hero","superhero"]
//输出：["as","hero"]
//解释："as" 是 "mass" 的子字符串，"hero" 是 "superhero" 的子字符串。
//["hero","as"] 也是有效的答案。
//输入：words = ["leetcode","et","code"]
//输出：["et","code"]
//解释："et" 和 "code" 都是 "leetcode" 的子字符串。
//输入：words = ["blue","green","bu"]
//输出：[]
//提示：
//1 <= words.length <= 100
//1 <= words[i].length <= 30
//words[i] 仅包含小写英文字母。
//题目数据 保证 每个 words[i] 都是独一无二的。

func stringMatching(words []string) []string {
	var ret []string
	for k1, v1 := range words {
		for k2, v2 := range words {
			if k1 != k2 && strings.Contains(v2, v1) {
				ret = append(ret, v1)
				break
			}
		}
	}
	return ret
}
