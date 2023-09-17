package d27

// gcd(nums[i],nums[j]) > 1，代表了两个数值拥有公共的因数，而每个因数最终都能分解为质因数
/*var MAX int
var ranges [][]int

// 初始化代码中做质因数分解，太久没写质因数分解代码，已经忘了怎么写了
func init() {
	MAX = 100000
	ranges = make([][]int, MAX+1)
	ranges[2] = make([]int, 1)
	ranges[2][0] = 2
	ranges[3] = make([]int, 1)
	ranges[3][0] = 3
	for i := 2; i <= MAX; i++ {
		x := i
		for j := 2; j*j <= x; j++ {
			if x%j == 0 {
				if ranges[i] == nil {
					ranges[i] = make([]int, 1)
					ranges[i][0] = j
				} else {
					ranges[i] = append(ranges[i], j)
				}
				for x%j == 0 {
					x /= j
				}
			}
		}
		if x > 1 {
			ranges[i] = append(ranges[i], x)
		}
	}
}

var parent []int

func findParent(i int) int {
	if parent[i] == i {
		return i
	}
	parent[i] = findParent(parent[i])
	return parent[i]
}
func canTraverseAllPairs(nums []int) bool {
	n := len(nums)
	parent = make([]int, n+MAX+1)
	for i := 0; i <= n+MAX; i++ {
		parent[i] = i
	}
	for i := 0; i < n; i++ {
		tmp := ranges[nums[i]]
		l := len(tmp)
		for j := 0; j < l; j++ {
			x := findParent(i)
			y := findParent(n + tmp[j])
			//fmt.Print("\nx = ",x)
			//fmt.Print("\ny = ",y)
			//fmt.Print("\ni = ",i)
			if x == y {
				continue
			}
			parent[x] = y
		}
	}
	pre := findParent(0)
	//fmt.Print("pre = ",pre)
	res := true
	for i := 1; i < n; i++ {
		now := findParent(i)
		if pre != now {
			res = false
			break
		}
	}
	return res
}*/

// 重新写一遍
// 解题思路是质因数分解+并查集
var MAX int
var ranges [][]int

func init() {
	MAX = 100000
	ranges = make([][]int, MAX)
	//初始化
	for i := 2; i < 4; i++ {
		ranges[i] = make([]int, 1)
		ranges[i][0] = i
	}
	for i := 4; i <= MAX; i++ {
		x := i
		for j := 2; j <= x/j; j++ {
			if x%j == 0 {
				if ranges[i] == nil {
					ranges[i] = make([]int, 1)
					ranges[i][0] = j
				} else {
					ranges[i] = append(ranges[i], j)
				}
				for x%j == 0 {
					x /= j
				}
			}
		}
		if x > 1 {
			ranges[i] = append(ranges[i], x)
		}
	}
}
func canTraverseAllPairs(nums []int) bool {
	n := len(nums)
	parent = make([]int, n+MAX+1)
	for i := 0; i < n+MAX+1; i++ {
		parent[i] = i
	}
	for i := 0; i < n; i++ {
		tmp := ranges[nums[i]]
		for j, _ := range tmp {
			x := findParent(i)
			y := findParent(n + j)
			if x == y {
				continue
			}
			parent[x] = y
		}
	}
	pre := findParent(0)
	for i := 1; i < n; i++ {
		now := findParent(i)
		if pre != now {
			return false
		}
	}
	return true
}

var parent []int

func findParent(loc int) int {
	if loc == parent[loc] {
		return loc
	}
	parent[loc] = findParent(parent[loc])
	return parent[loc]
}
