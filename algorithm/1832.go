package algorithm

func CheckIfPangram(sentence string) bool {
	memo := [26]int{}

	for _, c := range sentence {
		memo[c-'a']++
	}

	for _, v := range memo {
		if v == 0 {
			return false
		}
	}

	return true
}
