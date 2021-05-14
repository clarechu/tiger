package leet_code

import (
	"math"
)

func maxSubArray(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}
	if length < 2 {
		return nums[0]
	}
	max := nums[0]
	dp := make([]int, 0)
	dp[0] = nums[0]
	for l := 1; l < length; l++ {
		dp[l] = int(math.Max(float64(dp[l-1]+nums[l]), float64(nums[l])))
		if dp[l] > max {
			max = dp[l]
		}
	}
	return max
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	q := 1
	p := 2
	u := 1
	for i := 3; i <= n; i++ {
		u = q + p
		q = p
		p = u
	}
	return u
}

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	n := make([]int, len(nums)-1)
	for i, j := range nums {
		if i < 1 {
			n[i] = j
			continue
		}
		n[i] = n[i-1] + j
	}
	return NumArray{
		sum: n,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sum[right]
	}
	return this.sum[right] - this.sum[left-1]
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	return trap1(0, 0, height)
}

func trap1(cur, max int, height []int) int {
	flag := false
	if cur >= len(height) {
		return max
	}
	if height[cur] == 0 {
		return trap1(cur+1, max, height)
	}
	for i := cur + 2; i < len(height); i++ {
		if height[cur] <= height[i] {
			flag = true
			for j := cur + 1; j < i; j++ {
				max += height[cur] - height[j]
			}
			return trap1(i, max, height)
		}
	}
	if !flag {
		if cur == len(height)-1 {
			return max
		}
		return trap1(cur+1, max, height)
	}
	return max
}
func uniquePaths(m, n int) int {
	a := make([][]int, m)
	for i := 0; i < len(a); i++ {
		a[i] = make([]int, n)
		a[i][0] = 1
	}
	for j := 0; j < n; j++ {
		a[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			a[i][j] = a[i-1][j] + a[i][j-1]
		}
	}
	return a[m-1][n-1]
}

func minPathSum(grid [][]int) int {
	a := make([][]int, len(grid))
	for i := 0; i < len(a); i++ {
		a[i] = make([]int, len(grid[0]))
		if i == 0 {
			a[i][0] = grid[0][0]
		} else {
			a[i][0] = a[i-1][0] + grid[i][0]
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		if i == 0 {
			a[0][i] = grid[0][0]
		} else {
			a[0][i] = a[0][i-1] + grid[0][i]
		}
	}
	for i := 1; i < len(a); i++ {
		for j := 1; j < len(a[0]); j++ {
			if a[i-1][j] < a[i][j-1] {
				a[i][j] = a[i-1][j] + grid[i][j]
			} else {
				a[i][j] = a[i][j-1] + grid[i][j]
			}
		}
	}
	return a[len(grid)-1][len(grid[0])-1]
}
