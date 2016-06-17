package dynamic

func KnapsackDynamic(weights []int, values []int, n int, W int) [][]int {
	m := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		m[i] = make([]int, W+1)
	}
	for i := 0; i <= W; i++ {
		m[0][i] = 0
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= W; j++ {
			if weights[i] > j {
				m[i][j] = m[i-1][j]
			} else if m[i-1][j] > m[i-1][j-weights[i]]+values[i] {
				m[i][j] = m[i-1][j]
			} else {
				m[i][j] = m[i-1][j-weights[i]] + values[i]
			}
		}
	}
	return m
}
