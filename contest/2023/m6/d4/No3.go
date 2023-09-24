package d4

type pair struct{ val, time int }

// 使用时间戳的解法，刚好能卡过
/*func matrixSumQueries(n int, queries [][]int) int64 {
	rowCnt := make([]pair, n)
	colCnt := make([]pair, n)
	nowTime := 1
	for _, node := range queries {
		if node[0] == 0 {
			rowCnt[node[1]].val = node[2]
			rowCnt[node[1]].time = nowTime
		} else {
			colCnt[node[1]].val = node[2]
			colCnt[node[1]].time = nowTime
		}
		nowTime++
	}
	var res int64
	res = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += getNew(rowCnt[i], colCnt[j])
		}
	}
	return res
}

func getNew(x, y pair) int64 {
	if x.time > y.time {
		return int64(x.val)
	}
	return int64(y.val)
}*/

// 真正优秀的解法：逆向思维
func matrixSumQueries(n int, queries [][]int) int64 {
	rowMap, colMap := map[int]bool{}, map[int]bool{}
	m := len(queries)
	var res int64
	res = 0
	for i := m - 1; i >= 0; i-- {
		if queries[i][0] == 0 {
			if rowMap[queries[i][1]] {
				continue
			}
			rowMap[queries[i][1]] = true
			l := n - len(colMap)
			res += int64(l) * int64(queries[i][2])
		} else {
			if colMap[queries[i][1]] {
				continue
			}
			colMap[queries[i][1]] = true
			l := n - len(rowMap)
			res += int64(l) * int64(queries[i][2])
		}
	}
	return res
}
