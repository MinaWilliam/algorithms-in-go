package divideconquer

func Count(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return Count(nums[1:]) + 1
}
