package sort

func SelectionSort(nums []int) []int {
	len := len(nums)
	for i := 0; i < len; i++ {
		smallestIndex := findSmallest(nums[i:])
		nums[i], nums[smallestIndex+i] = nums[smallestIndex+i], nums[i]
	}
	return nums
}

func findSmallest(nums []int) int {
	smallest := nums[0]
	smallestIndex := 0
	for i, v := range nums {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}
	return smallestIndex
}
