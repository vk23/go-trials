package leetcode

import (
	"testing"
)

var count = make(map[int]int)

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if v, ok := count[n]; ok {
		return v
	}
	count[n] = climbStairs(n-1) + climbStairs(n-2)
	return count[n]
}

func TestClimbStairs(t *testing.T) {
	type testCase struct {
		stairs   int
		expected int
	}
	cases := []testCase{
		{2, 2},
		{3, 3},
	}

	for _, cs := range cases {
		res := climbStairs(cs.stairs)
		if res != cs.expected {
			t.Errorf("climbStairs(%v) == %v, expected %v", cs.stairs, res, cs.expected)
		}
	}
}
