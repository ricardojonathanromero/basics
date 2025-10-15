package two_nums

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(left *ListNode, right *ListNode) *ListNode {
	root := &ListNode{}
	current := root

	carry := 0
	for left != nil || right != nil {
		leftVal, rightVal := 0, 0
		if left != nil {
			leftVal = left.Val
			left = left.Next
		}
		if right != nil {
			rightVal = right.Val
			right = right.Next
		}

		sum := leftVal + rightVal + carry
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10, Next: nil}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return root.Next
}
