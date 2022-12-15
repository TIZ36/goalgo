package algorithm

func GetLucky(s string, k int) int {

	var sum int32
	var buf []int32

	// "adfdaf"
	for _, c := range s {
		diff := c - 'a' + 1
		if diff < 10 {
			buf = append(buf, diff)
		} else {
			buf = append(buf, diff/10)
			buf = append(buf, diff%10)
		}
	}

	for ; k > 0; k-- {
		sum = getSum(buf)
		buf = splitNum(sum)
	}

	return int(sum)
}

func getSum(buf []int32) (sum int32) {
	for i := 0; i < len(buf); i++ {
		sum += buf[i]
	}

	return
}

func splitNum(num int32) []int32 {
	var buf []int32
	for num > 0 {
		buf = append(buf, num%10)
		num /= 10
	}

	swapList(buf)

	return buf
}

func swapList(list []int32) {
	start := 0
	end := len(list) - 1

	for start < end {
		list[start], list[end] = list[end], list[start]
		start++
		end--
	}
}
