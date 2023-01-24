package easy

func containsDuplicate(nums []int) bool {
	dups := make(map[int]int)
	for i := range nums {
		if _, ok := dups[nums[i]]; ok {
			return true
		}
		dups[nums[i]] = nums[i]
	}
	return false
}
