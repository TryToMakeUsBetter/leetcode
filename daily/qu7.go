package daily

// lc121
func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][0] = max(dp[i-1][0], -prices[i])
	}

	return dp[len(prices)-1][1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
