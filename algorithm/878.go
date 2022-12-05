package algorithm

func NthMagicalNumber(n int, a int, b int) int {
	lcm := a / gcd(a, b) * b
	left := min(a, b)
	right := min(a, b) * n

	var mid int
	for left+1 < right {
		mid = (right-left)/2 + left

		num := mid/a + mid/b - mid/lcm

		if num >= n {

			right = mid
		} else {
			left = mid
		}
	}

	return (right) % (1e9 + 7)
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}

	return a
}
