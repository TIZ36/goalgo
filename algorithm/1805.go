package algorithm

import "unicode"

func NumDifferentIntegers(word string) int {
	s := map[string]bool{}
	n := len(word)
	p1 := 0
	for {
		for p1 < n && !unicode.IsDigit(rune(word[p1])) {
			p1++
		}
		if p1 == n {
			break
		}
		p2 := p1
		for p2 < n && unicode.IsDigit(rune(word[p2])) {
			p2++
		}
		for p2-p1 > 1 && word[p1] == '0' { // 去除前导 0
			p1++
		}
		s[word[p1:p2]] = true
		p1 = p2
	}
	return len(s)
}

// 对超级大数没办法
//func NumDifferentIntegers(word string) int {
//	memo := make(map[int32]int)
//	var acc int32 = -1
//
//	for _, c := range word {
//		if c >= '0' && c <= '9' {
//			if acc == -1 {
//				acc = 0
//			}
//
//			acc = acc*10 + (c - '0')
//		} else {
//			memo[acc] = 1
//			acc = -1
//		}
//	}
//
//	if acc != -1 {
//		memo[acc] = 1
//	}
//
//	delete(memo, -1)
//
//	return len(memo)
//}
