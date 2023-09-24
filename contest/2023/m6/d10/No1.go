package d10

import "strconv"

func isFascinating(n int) bool {
	n2 := 2 * n
	n3 := 3 * n
	s := strconv.Itoa(n) + strconv.Itoa(n2) + strconv.Itoa(n3)
	vis := make([]bool, 10)
	for i := 0; i < len(s); i++ {
		num := int(s[i] - '0')
		if num == 0 || vis[num] {
			return false
		}
		vis[num] = true
	}
	return true
}
