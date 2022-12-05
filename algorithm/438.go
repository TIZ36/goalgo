package algorithm

import "fmt"

func FindAnagrams(s string, p string) (ans []int) {
	if len(s) < len(p) {
		return nil
	}

	primDic := []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

	target, temp := 1, 1
	plen := len(p)

	// get the target value of p
	for _, v := range p {
		target *= primDic[v-'a']
	}

	fmt.Println("target is :", target)

	for _, v := range s[:plen] {
		temp *= primDic[v-'a']
	}

	for idx, v := range s[plen:] {

		fmt.Println("temp is :", temp)
		fmt.Println("--- ---")
		if temp == target {
			ans = append(ans, idx)
		}

		temp = temp / primDic[s[idx]-'a'] * primDic[v-'a']
	}

	if temp == target {
		ans = append(ans, len(s)-plen)
	}

	return
}
