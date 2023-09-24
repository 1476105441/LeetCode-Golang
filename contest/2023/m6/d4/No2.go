package d4

func semiOrderedPermutation(nums []int) int {
	n := len(nums)
	res := 0
	if nums[0] != 1 {
		i := 0
		for ; i < n; i++ {
			if nums[i] == 1 {
				break
			}
		}
		res += i
		for ; i > 0; i-- {
			tmp := nums[i-1]
			nums[i-1] = nums[i]
			nums[i] = tmp
		}
	}
	if nums[n-1] != n {
		i := n - 1
		for ; i > 0; i-- {
			if nums[i] == n {
				break
			}
		}
		res += n - 1 - i
	}
	return res
}
