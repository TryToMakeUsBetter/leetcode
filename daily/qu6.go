package daily

// lc189
func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	if k > len(nums)-1 {
		k %= len(nums)
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
