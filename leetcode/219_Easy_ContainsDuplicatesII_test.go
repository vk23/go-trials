package leetcode

import (
	"testing"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool, k)
	for i, v := range nums {
		if i > k {
			delete(m, nums[i-k-1])
		}
		if m[v] {
			return true
		} else {
			m[v] = true
		}
	}
	return false
}

func TestContainsDuplicateII(t *testing.T) {
	type testCase struct {
		array    []int
		k        int
		expected bool
	}

	var cases = []testCase{
		{[]int{1, 2, 3, 1, 2, 3}, 3, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, cs := range cases {
		res := containsNearbyDuplicate(cs.array, cs.k)
		if res != cs.expected {
			t.Errorf("containsNearbyDuplicate(%v, %v) == %v, expected %v", cs.array, cs.k, res, cs.expected)
		}
	}
}
