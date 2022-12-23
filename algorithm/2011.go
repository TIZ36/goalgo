package algorithm

import "strings"

func FinalValueAfterOperations(operations []string) int {

	var re int

	for _, ss := range operations {
		if strings.Contains(ss, "++") {
			re++
		} else {
			re--
		}
	}

	return re
}
