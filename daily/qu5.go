package daily

// lc169
func majorityElement(nums []int) int {
	if len(nums) <= 2 {
		return nums[0]
	}

	cnt := 1
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if ans == nums[i] {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			ans = nums[i]
		}
	}

	return ans
}
