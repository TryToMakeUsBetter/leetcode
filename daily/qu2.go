package daily

// lc27
func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	tail := len(nums) - 1
	i := 0
	for ; i < tail; i++ {
		if nums[i] == val {
			nums[tail], nums[i] = nums[i], nums[tail]
			i--
			tail--
		}

	}

	if nums[i] == val {
		nums[tail], nums[i] = nums[i], nums[tail]
		i--
		tail--
	}

	return i + 1
}
