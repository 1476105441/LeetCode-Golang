package d28

func minimumCost(s string) int64 {
	n := len(s)
	right := make([]int, n)
	// 左边全部规整的代价
	var tmp int
	var pre uint8 = s[n-1]
	for i := n - 2; i >= 0; i-- {
		if s[i] != pre {
			right[i] = right[i+1] + n - i - 1
			pre = s[i]
		} else {
			right[i] = right[i+1]
		}
	}
	tmp = 0
	pre = s[0]
	res := right[0]
	for i := 1; i < n; i++ {
		if s[i] != pre {
			val := right[i] + n - i + tmp
			tmp = tmp + i
			if val > tmp+right[i] {
				val = tmp + right[i]
			}
			if res > val {
				res = val
			}
			pre = s[i]
		} else {
			if res > tmp+right[i] {
				res = tmp + right[i]
			}
		}
	}
	return int64(res)
}
