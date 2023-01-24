package easy

func findDisappearedNumbers(nums []int) []int {
	set := make(map[int]struct{}, len(nums))
	disappeared := make([]int, 0, len(nums)-2)

	for _, num := range nums {
		set[num] = struct{}{}
	}

	for i := 1; i <= len(nums); i++ {
		if _, exists := set[i]; !exists {
			disappeared = append(disappeared, i)
		}
	}

	if len(disappeared) != 0 {
		return disappeared
	}
	return []int{}
}
