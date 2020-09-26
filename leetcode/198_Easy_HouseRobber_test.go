package leetcode

import (
	"testing"
)

var cache map[int]int

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	cache = make(map[int]int, len(nums))
	return max(recur(nums, 0), recur(nums, 1))
}

func recur(nums []int, i int) int {
	if i >= len(nums) {
		return 0
	}
	if v, ok := cache[i]; ok {
		return v
	}
	val := max(recur(nums, i+2), recur(nums, i+3)) + nums[i]
	cache[i] = val
	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestHouseRobber(t *testing.T) {
	type testCase struct {
		houses   []int
		expected int
	}
	cases := []testCase{
		{[]int{2, 7, 9, 3, 1, 12, 3, 4, 5, 5}, 32},
		{[]int{2, 1, 1, 2}, 4},
		{[]int{1, 2, 3, 1}, 4},
	}

	for _, cs := range cases {
		res := rob(cs.houses)
		if res != cs.expected {
			t.Errorf("rob(%v) == %v, expected %v", cs.houses, res, cs.expected)
		}
	}
}
