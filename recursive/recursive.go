package recursive

func KnapsackRecursive(weights []int, values []int, n int, W int) int {
	if n == 0 || W == 0 {
		return 0
	}
	without := KnapsackRecursive(weights, values, n-1, W)
	if weights[n] > W {
		return without
	}
	withim := values[n] + KnapsackRecursive(weights, values, n-1, W-weights[n])
	if withim > without {
		return withim
	}
	return without
}
