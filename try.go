package main

func getMinimumTime(time []int, fruits [][]int, limit int) int {
	var cout, tm int
	for _, v := range fruits {
		if v[1]%limit == 0 {
			cout = v[1] / limit
		} else {
			cout = v[1]/limit + 1
		}
		tm += cout * time[v[0]]
		cout = 0
	}
	return tm
}
