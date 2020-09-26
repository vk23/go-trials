package leetcode

import (
	"testing"
)

/*
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
*/
func removeDuplicates(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}
	tmp, n := nums[0], 1
	for i := 1; i < size; i++ {
		if nums[i] == tmp {
			continue
		}
		nums[n], n, tmp = nums[i], n+1, nums[i]
	}
	return n
}

func TestRemoveDuplicatesFromSorted(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}
	cases := []testCase{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, cs := range cases {
		//fmt.Printf("before: %v\n", cs.nums)
		res := removeDuplicates(cs.nums)
		//fmt.Printf("after: %v\n", cs.nums)
		if res != cs.expected {
			t.Errorf("removeDuplicates(%v) == %v, expected %v", cs.nums, res, cs.expected)
		}
	}
}
