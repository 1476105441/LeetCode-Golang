package d10

import "sort"

func sumDistance(nums []int, s string, d int) int {
	for i, x := range s {
		if x == 'R' {
			nums[i] += d
		} else {
			nums[i] -= d
		}
	}
	sort.Ints(nums)
	sum := 0
	var mod int = 1e9 + 7
	res := 0
	for i, x := range nums {
		res = (res + i*x - sum) % mod
		sum += x
	}
	return res
}
