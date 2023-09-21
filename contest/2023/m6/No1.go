package m6

func minimizedStringLength(s string) int {
	runes := map[rune]int{}
	for _, c := range s {
		runes[c] = 1
	}
	return len(runes)
}
