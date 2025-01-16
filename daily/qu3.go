package daily

// lc26
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow := 0

	for fast := slow + 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			nums[slow+1] = nums[fast]
			slow++

		}
	}
	return slow + 1
}
