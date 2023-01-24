package easy

func singleNumber(nums []int) int {
	var mask int

	for _, num := range nums {
		mask ^= num
	}

	return mask
}
