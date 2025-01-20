package daily

import "math/rand"

// lc380
type RandomizedSet struct {
	nums []int
	mem  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		mem:  map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mem[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	pos := len(this.nums) - 1
	this.mem[val] = pos

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	pos, ok := this.mem[val]
	if !ok {
		return false
	}

	lastPos := len(this.nums) - 1
	this.nums[pos] = this.nums[lastPos]
	this.mem[this.nums[pos]] = pos
	this.nums = append(this.nums, val)
	this.nums = this.nums[:lastPos]
	delete(this.mem, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
