package algorithm

func Gcd(a, b int) int {
	if b != 0 {
		return Gcd(b, a%b)
	}

	return a
}
