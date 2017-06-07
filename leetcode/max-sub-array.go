package leetcode

func maxSubArray(nums []int) int {
	length := len(nums)

	dp := make([]int, length)
	dp[0] = nums[0]
	max := dp[0]

	for i := 1 ; i < length; i++{

		var previous int
		if dp[i-1] > 0 {
			previous = dp[i-1]
		}else{
			previous = 0
		}
		dp[i] = nums[i] + previous
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
