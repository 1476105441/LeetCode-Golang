package d28

func removeTrailingZeros(num string) string {
	n := len(num)
	loc := n
	for i := n - 1; i >= 0; i-- {
		if num[i] == '0' {
			loc--
		} else {
			break
		}
	}
	return num[:loc]
}
