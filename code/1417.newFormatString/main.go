package main

import (
	"fmt"
)

func main() {
	s := "1231fasbd"
	fmt.Println(reformat(s))
}

func reformat1(s string) string {
	s1, s2 := "", ""
	for _, v := range s {
		if v >= 48 && v <= 57 {
			s1 += string(v)
			continue
		}
		s2 += string(v)
	}
	if len(s1)-len(s2) <= 1 && len(s1)-len(s2) >= -1 {
		if len(s1) < len(s2) {
			s1, s2 = s2, s1
		}
		ret := ""
		for i, _ := range s2 {
			ret += string(s1[i]) + string(s2[i])
		}
		if len(s1) != len(s2) {
			ret += string(s1[len(s1)-1])
		}
		return ret
	}
	return ""
}

func reformat(s string) string {
	ret, c1, c2 := "", []string{}, []string{}
	for _, v := range s {
		if v >= 48 && v <= 57 {
			c1 = append(c1, string(v))
			continue
		}
		c2 = append(c2, string(v))
	}
	sub := len(c1) - len(c2)
	if sub <= 1 && sub >= -1 {
		if sub < 0 {
			c1, c2 = c2, c1
		}
		for i := range c2 {
			ret += c1[i] + c2[i]
		}
		if sub != 0 {
			ret += c1[len(c1)-1]
		}
	}
	return ret
}
