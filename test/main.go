package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
func maxSlidingWindow(nums []int, k int) []int {

	left, right, n := 0, k, len(nums)
	ans := []int{}
	sum, maxSum := 0, 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum, ans = sum, append(ans, sum)
	for right < n {
		sum += nums[right]
		right++
		sum -= nums[left]
		left++
		maxSum = max(maxSum, sum)
		ans = append(ans, maxSum)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
