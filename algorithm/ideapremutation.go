package algorithm

func IsIdealPermutation(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return true
	}

	length := len(nums)
	premax := nums[0]

	for i := 2; i < length; i++ {
		if nums[i] < premax {
			return false
		}

		premax = max(nums[i-1], premax)
	}

	return true
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
