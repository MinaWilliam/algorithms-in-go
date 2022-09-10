package divideconquer

func Sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return Sum(nums[1:]) + nums[0]
}
