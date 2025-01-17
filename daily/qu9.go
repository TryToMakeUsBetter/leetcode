package daily

// lc55
func canJump(nums []int) bool {
	maxGet := 0
	l := len(nums)
	for i := 0; i < len(nums); i++ {
		if maxGet >= l-1 {
			return true
		}
		if i < maxGet {
			maxGet = max(maxGet, i+nums[i])
		}
	}
	if maxGet >= l-1 {
		return true
	}

	return false
}
