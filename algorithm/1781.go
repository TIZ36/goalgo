package algorithm

import (
	"math"
)

func BeautySum(s string) int {

	length := len(s)

	if length < 3 {
		return 0
	}

	buf := map[string]int{}

	for i := 0; i < length-2; i++ {
		for z := i + 2; z < length; z++ {
			txt := s[i : z+1]
			buf[txt] += BeautyDiff(txt)
		}
	}

	var re int
	for _, v := range buf {
		re += v
	}

	return re
}

func BeautyDiff(s string) int {
	if len(s) < 3 {
		return 0
	}

	memo := make([]int, 26)
	maxx, minn := math.MinInt, math.MaxInt

	for _, c := range s {
		memo[c-'a']++
	}

	for _, freq := range memo {
		if freq > 0 {
			maxx = mmax(freq, maxx)
			minn = mmin(freq, minn)
		}
	}

	return mmax(maxx-minn, 0)
}

func mmax(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func mmin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
