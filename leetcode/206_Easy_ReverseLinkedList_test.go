package leetcode

import (
	"testing"
)

// iteratively
func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// recursively
func reverseList2(head *ListNode) *ListNode {
	var recur func(*ListNode, *ListNode) *ListNode
	recur = func(prev, cur *ListNode) *ListNode {
		if cur == nil {
			return prev
		}
		next := cur.Next
		cur.Next = prev

		return recur(cur, next)
	}

	return recur(nil, head)
}

func toLinkedList(source []int) *ListNode {
	var root = new(ListNode)
	node := root
	for _, v := range source {
		node.Next = &ListNode{v, nil}
		node = node.Next
	}

	return root.Next
}

func TestReverseLinkedList(t *testing.T) {
	type testCase struct {
		original []int
		expected []int
	}

	var cases = []testCase{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1}, []int{1}},
	}

	for _, cs := range cases {
		expected := toLinkedList(cs.expected)
		res := reverseList1(toLinkedList(cs.original))
		if !equal(res, expected) {
			t.Errorf("reverseList1(%v) == %v, expected %v", cs.original, res, cs.expected)
		}
		res = reverseList2(toLinkedList(cs.original))
		if !equal(res, expected) {
			t.Errorf("reverseList2(%v) == %v, expected %v", cs.original, res, cs.expected)
		}
	}
}
