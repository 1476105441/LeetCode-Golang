package d27

import "sort"

// 想法：大于0的都可以选，负数则需要看情况选
func maxStrength(nums []int) int64 {
	n := len(nums)
	if n == 1 {
		return int64(nums[0])
	}
	negative := make([]int, n)
	l := 0
	res := int64(1)
	flag := false
	for _, num := range nums {
		if num > 0 {
			res *= int64(num)
			flag = true
		} else if num < 0 {
			negative[l] = num
			l++
		}
	}
	sort.Ints(negative)
	var tmp []int
	if l&1 == 1 {
		l--
	}
	tmp = negative[:l]
	for _, x := range tmp {
		res *= int64(x)
		flag = true
	}
	if flag {
		return res
	} else {
		return 0
	}
}
