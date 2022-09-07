package search

func BinarySearch(nums []int, n int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2
		guess := nums[mid]
		if n == guess {
			return mid
		}
		if n > guess {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func BinaryExist(nums []int, n int) bool {
	mid := int(len(nums) / 2)

	if n == nums[mid] {
		return true
	}
	if n > nums[mid] && mid > 0 {
		return BinaryExist(nums[mid:], n)
	}
	if n < nums[mid] && mid > 0 {
		return BinaryExist(nums[:mid], n)
	}
	return false
}
