package leetcode

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	} else if right == nil {
		return left
	}

	var res *ListNode
	if left.Val <= right.Val {
		res, left = left, left.Next
	} else {
		res, right = right, right.Next
	}

	cur := res
	for {
		if left == nil && right == nil {
			break
		} else if left == nil {
			cur.Next = right
			break
		} else if right == nil {
			cur.Next = left
			break
		}
		if left.Val <= right.Val {
			cur.Next, left = left, left.Next
		} else {
			cur.Next, right = right, right.Next
		}
		cur = cur.Next
	}
	return res
}

func TestMergeTwoLists(t *testing.T) {
	type testCase struct {
		l1, l2   *ListNode
		expected *ListNode
	}
	cases := []testCase{
		{
			l1: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 4, Next: nil,
					},
				},
			},

			l2: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 4, Next: nil,
					},
				},
			},

			expected: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 3, Next: &ListNode{
								Val: 4, Next: &ListNode{
									Val: 4, Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, cs := range cases {
		res := mergeTwoLists(cs.l1, cs.l2)
		print(cs.l1)
		print(cs.l2)
		print(res)
		if !equal(res, cs.expected) {
			t.Errorf("mergeTwoLists(%+v, %+v) == %+v, expected %+v", cs.l1, cs.l2, res, cs.expected)
		}
	}
}

func equal(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil || l2 == nil {
		return false
	}
	for {
		if (l1 == nil) != (l2 == nil) {
			return false
		} else if l1 == nil && l2 == nil {
			break
		}
		if l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
	return true
}

func print(l *ListNode) {
	if l == nil {
		fmt.Println("List: nil")
	}
	for l != nil {
		fmt.Printf("%v->", l.Val)
		l = l.Next
	}
	fmt.Println()
}
