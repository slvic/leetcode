package easy

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, 0)

	currentSum := 0

	for _, num := range nums {
		currentSum = currentSum + num
		sums = append(sums, currentSum)
	}

	var numArray NumArray

	numArray.sums = sums

	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	}
	return this.sums[right] - this.sums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
