package sort

func QuickSort(nums []int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return nums
	}

	// todo: pick an in middle pivot to avoid O(n^2) if the array is sorted.
	pivot := nums[0]
	greater := make([]int, 0)
	less := make([]int, 0)

	for i := 1; i < numsLen; i++ {
		if nums[i] > pivot {
			greater = append(greater, nums[i])
		} else {
			less = append(less, nums[i])
		}
	}

	return append(
		append(QuickSort(less), pivot),
		QuickSort(greater)...,
	)
}
