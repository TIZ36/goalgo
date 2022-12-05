package algorithm

import "fmt"

type Cnum struct {
	charc int32
	num   int32
}

func ExpressiveWords(s string, words []string) int {
	// 1、词序一致
	// 2、同位词，个数相等或 >= 3

	var re int

	stringCSlice := GetCnum(s)

	fmt.Println(stringCSlice)

	for _, word := range words {

		isextend := true
		wCSlice := GetCnum(word)

		fmt.Println(wCSlice)

		if len(wCSlice) != len(stringCSlice) {

			isextend = false
			continue
		} else {
			for idx, wcum := range wCSlice {
				scum := stringCSlice[idx]

				if wcum.charc != scum.charc || (wcum.num != scum.num && scum.num < 3) || wcum.
					num > scum.num {
					isextend = false
					break
				}
			}
		}

		if isextend {
			re++
		}
	}

	return re
}

func GetCnum(s string) (cnumSlice []Cnum) {
	var curc int32
	for _, c := range s {
		if c == curc {
			cnumSlice[len(cnumSlice)-1].num++
		} else {
			curc = c
			newc := Cnum{
				charc: c,
				num:   1,
			}

			cnumSlice = append(cnumSlice, newc)
		}
	}

	return
}
