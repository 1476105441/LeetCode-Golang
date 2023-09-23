package m6

import "strconv"

// 注意golang中指针的使用
var s string

func numDupDigitsAtMostN(n int) int {
	s = strconv.Itoa(n)
	m := len(s)
	vis := make([][1 << 10]int, m)
	for i := range vis {
		for j := range vis[i] {
			vis[i][j] = -1
		}
	}
	return n - f(0, 0, true, false, &vis)
}
func f(loc, mask int, isLimit, isNum bool, vis *[][1 << 10]int) int {
	// golang不支持三目运算符
	if loc == len(s) {
		if isNum {
			return 1
		}
		return 0
	}
	if !isLimit && isNum && (*vis)[loc][mask] != -1 {
		return (*vis)[loc][mask]
	}
	res := 0
	if !isNum {
		res = f(loc+1, mask, false, false, vis)
	}
	var top int
	var begin int
	if isLimit {
		top = int(s[loc] - '0')
	} else {
		top = 9
	}
	if isNum {
		begin = 0
	} else {
		begin = 1
	}
	for i := begin; i <= top; i++ {
		if (mask >> i & 1) == 0 {
			res += f(loc+1, mask|(1<<i), isLimit && (i == top), true, vis)
		}
	}
	if !isLimit && isNum {
		(*vis)[loc][mask] = res
	}
	return res
}
