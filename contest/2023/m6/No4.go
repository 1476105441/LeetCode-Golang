package m6

//
var str string
var min, max int
var mod int

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	mod = 1e9 + 7
	min = min_sum
	max = max_sum
	res := getVal(num2) - getVal(num1) + mod
	c := 0
	for _, cs := range num1 {
		c += int(cs - '0')
	}
	if min_sum <= c && c <= max_sum {
		res = (res + 1) % mod
	}
	return res % mod
}

func getVal(num string) int {
	str = num
	m := len(num)
	vis := make([][1 << 10]int, m)
	for i := range vis {
		for j := range vis[i] {
			vis[i][j] = -1
		}
	}
	return fn(0, 0, true, false, &vis)
}

func fn(loc, mask int, isLimit, isNum bool, vis *[][1 << 10]int) int {
	if mask > max {
		return 0
	}
	if loc == len(str) {
		if min <= mask {
			return 1
		}
		return 0
	}
	if !isLimit && (*vis)[loc][mask] != -1 {
		return (*vis)[loc][mask]
	}
	res := 0
	var top int
	if isLimit {
		top = int(s[loc] - '0')
	} else {
		top = 9
	}
	for i := 0; i <= top; i++ {
		res = (res + f(loc+1, mask+i, isLimit && i == top, false, vis)%mod) % mod
	}
	if !isLimit {
		(*vis)[loc][mask] = res
	}
	return res
}
