package daily

// lc45
func jump(nums []int) int {
	maxPos := 0
	length := len(nums)
	step := 0
	end := 0
	for i := 0; i < length-1; i++ {
		maxPos = max(maxPos, nums[i]+i)
		if i == end {
			step++
			end = maxPos
		}
	}
	return step
}
