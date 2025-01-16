package daily

// lc80
func removeDuplicatesV2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slow := 0
	maxShow := 1
	for fast := 1; fast < len(nums); fast++ {
		if maxShow != 0 && nums[slow] == nums[fast] {
			maxShow--
			nums[slow+1] = nums[fast]
			slow++
			continue
		}
		if nums[slow] != nums[fast] {
			maxShow = 1
			nums[slow+1] = nums[fast]
			slow++
		}
	}

	return slow + 1
}
