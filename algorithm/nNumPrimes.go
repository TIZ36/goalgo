package algorithm

import "math"

func GenPrimNum(n int, nums int) []int {

	if n == 1 {
		return []int{}
	}

	if n == 2 {
		return []int{2}
	}

	var re []int
	var isPrim bool
	for i := 3; i <= n; i++ {
		isPrim = true
		sqrtNum := int(math.Sqrt(float64(i)))
		for j := 2; j <= sqrtNum; j++ {
			if i%j == 0 {
				isPrim = false
				break
			}
		}

		if isPrim {
			re = append(re, i)
		}

		if len(re) == nums {
			return re
		}
	}

	return re
}
