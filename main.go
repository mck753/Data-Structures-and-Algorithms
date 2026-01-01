package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	copy(dp, triangle[n-1])
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}

	return dp[0]
}

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
	}

	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < len(grid[i]); j++ {
			x, y := i-1, j-1
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if x < 0 {
				dp[i][j] = grid[i][j] + dp[i][y]
			} else if y < 0 {
				dp[i][j] = grid[i][j] + dp[x][j]
			} else {
				dp[i][j] = grid[i][j] + min(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[n-1][len(grid[0])-1]
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, p := len(s1), len(s2), len(s3)
	if m+n != p {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	for j := 1; j <= n; j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i <= m; i++ {
		if s1[i-1] == s2[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j < n; j++ {
			if s3[i+j-1] == s1[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if s3[i+j-1] == s2[j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}

	}

	return dp[n][p]
}

func main() {
	triangle := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(triangle))
}
