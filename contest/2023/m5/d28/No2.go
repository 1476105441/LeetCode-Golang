package d28

func differenceOfDistinctValues(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])
	var maps [][]map[int]bool
	maps = make([][]map[int]bool, m)
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		maps[i] = make([]map[int]bool, n)
		maps[i][0] = make(map[int]bool)
		res[i] = make([]int, n)
	}
	for i := 1; i < n; i++ {
		maps[0][i] = make(map[int]bool)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				res[i][j] = len(maps[i-1][j-1])
				maps[i][j] = maps[i-1][j-1]
			} else {
				res[i][j] = 0
			}
			maps[i][j][grid[i][j]] = true
		}
	}
	for i := m - 1; i > -1; i-- {
		maps[i][n-1] = make(map[int]bool)
	}
	for i := n - 1; i > -1; i-- {
		maps[m-1][i] = make(map[int]bool)
	}

	for i := m - 1; i > -1; i-- {
		for j := n - 1; j > -1; j-- {
			var tmp int
			if i < m-1 && j < n-1 {
				tmp = len(maps[i+1][j+1])
				maps[i][j] = maps[i+1][j+1]
			} else {
				tmp = 0
			}
			maps[i][j][grid[i][j]] = true
			if res[i][j] > tmp {
				res[i][j] -= tmp
			} else {
				res[i][j] = tmp - res[i][j]
			}
		}
	}
	return res
}
