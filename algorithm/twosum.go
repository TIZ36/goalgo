package algorithm

import (
	"github.com/tiz36/goalgo/datastructure"
)

func addTwoNumbers(l1 *datastructure.ListNode, l2 *datastructure.ListNode) *datastructure.ListNode {
	carry := 0

	re := &datastructure.ListNode{}
	pp := re
	for l1 != nil && l2 != nil {
		total := (l1.Val + l2.Val) + carry
		sum := total % 10
		carry = total / 10

		re.Next = &datastructure.ListNode{Val: sum}
		re = re.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		total := l1.Val + carry
		sum := total % 10
		carry = total / 10

		re.Next = &datastructure.ListNode{Val: sum}
		re = re.Next

		l1 = l1.Next
	}

	for l2 != nil {
		total := l2.Val + carry
		sum := total % 10
		carry = total / 10

		re.Next = &datastructure.ListNode{Val: sum}
		re = re.Next

		l2 = l2.Next
	}

	return pp.Next
}
