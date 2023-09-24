package d10

func longestSemiRepetitiveSubstring(s string) int {
	n := len(s)
	l, r, cnt := 0, 0, 0
	pre := byte('a')
	res := 0
	for r < n {
		if s[r] == pre {
			if cnt == 0 {
				cnt++
			} else {
				// 需要滑动左窗口了，消除掉一个重复的串
				res = max(res, r-l)
				for l < r {
					l++
					if s[l-1] == s[l] {
						break
					}
				}
				// cnt不需要变化，因为要加上当前下标的字符
			}
		}
		pre = s[r]
		r++
	}
	if l < r {
		res = max(res, r-l)
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
