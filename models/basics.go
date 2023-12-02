package models

// 最长递增子序列1
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	max := 1
	for i:=1; i<n; i++ {
		dp[i] = 1
		for j:=0; j<i; j++ {
			if nums[i]>nums[j] {
				dp[i] = maxv(dp[i], dp[j]+1)
			}
		}
		max = maxv(max, dp[i])
	}
	return max
}

func maxv(i, j int) int {
	if i>j {
		return i
	}
	return j
}

// 最长递增子序列2
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	max := 1
	for i:=1; i<n; i++ {
		inx := biSearch(dp, max, nums[i])
		dp[inx] = nums[i]
		if inx == max {
			max++
		}
	}
	return max
}

func biSearch(nums []int, right, target int) int {
	left := 0
	for left<right {
		mid := (left+right)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid+1
		}
	}
	return left
}
