package algorithm

func CountHomogenous(s string) int {
	var mod int = 1e9 + 7

	cnt := 0
	pre := int32(s[0])
	var re int = 0

	for _, c := range s {
		if c == pre {
			cnt++
		} else {
			re += cnt * (cnt + 1) / 2
			cnt = 1
			pre = c
		}
	}

	re += cnt * (cnt + 1) / 2

	return re % mod
}
