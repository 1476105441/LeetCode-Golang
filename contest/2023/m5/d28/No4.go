package d28

import "sort"

// 实际上就是按照值从小到大来更新状态数组
func maxIncreasingCells(mat [][]int) int {
	type pair struct{ x, y int }
	m := len(mat)
	n := len(mat[0])
	rowMax := make([]int, m)
	colMax := make([]int, n)
	maps := map[int][]pair{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maps[mat[i][j]] = append(maps[mat[i][j]], pair{i, j})
		}
	}
	rank := make([]int, 0, len(maps))
	for key := range maps {
		rank = append(rank, key)
	}
	sort.Ints(rank)
	res := 0
	for _, key := range rank {
		nodes := maps[key]
		mx := make([]int, len(nodes))
		for i, node := range nodes {
			mx[i] = max(rowMax[node.x], colMax[node.y]) + 1
			res = max(res, mx[i])
		}
		for i, node := range nodes {
			rowMax[node.x] = max(rowMax[node.x], mx[i])
			colMax[node.y] = max(colMax[node.y], mx[i])
		}
	}
	return res
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
