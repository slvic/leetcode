package easy

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}

	sort.SliceStable(nums1, func(i, j int) bool {
		return nums1[i] <= nums1[j]
	})
}
