package two_nums

import "testing"

func TestTwoNumbers(t *testing.T) {
	var tests = []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			name: "first",
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "simple",
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
			l2: &ListNode{
				Val:  0,
				Next: nil,
			},
			expected: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := addTwoNumbers(test.l1, test.l2)
			for result.Next != nil {
				t.Log(result.Val)
				result = result.Next
			}
			t.Log(result.Val)
		})
	}
}
