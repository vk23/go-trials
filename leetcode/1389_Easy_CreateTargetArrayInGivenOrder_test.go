package leetcode

import (
	"fmt"
	"testing"
)

func createTargetArray(nums []int, indices []int) []int {
	res := make([]int, 1)
	for i, v := range nums {
		index := indices[i]
		res = append(res[0:index+1], res[index:]...)
		res[index] = v
	}
	return res[:len(nums)]
}

func TestTargetArrayInGivenOrder(t *testing.T) {
	type testCase struct {
		nums, indices []int
		expected      []int
	}
	cases := []testCase{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}, []int{0, 4, 1, 3, 2}},
		{[]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}, []int{0, 1, 2, 3, 4}},
		{[]int{1}, []int{0}, []int{1}},
	}

	for _, cs := range cases {
		fmt.Printf("before: %v\n", cs.nums)
		res := createTargetArray(cs.nums, cs.indices)
		if !equalArrays2(res, cs.expected) {
			t.Errorf("createTargetArray(%v, %v) == %v, expected %v", cs.nums, cs.indices, res, cs.expected)
		}
	}
}

func equalArrays2(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}
