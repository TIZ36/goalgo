package algorithm

func NumMatchingSubseq(s string, words []string) (ans int) {
	dic := [26][]string{}

	for _, v := range words {
		dic[v[0]-'a'] = append(dic[v[0]-'a'], v)
	}

	result := 0

	for _, c := range s {
		temp := dic[c-'a']
		dic[c-'a'] = nil

		for _, v := range temp {
			if len(v) == 1 {
				result++
			} else {
				dic[v[1]-'a'] = append(dic[v[1]-'a'], v[1:])
			}
		}
	}

	return result
}
