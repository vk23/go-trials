package leetcode

import (
	"testing"
)

/*
Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
Return the array in the form [x1,y1,x2,y2,...,xn,yn].
In-place?
*/
func shuffle(nums []int, n int) []int {
	sz := len(nums)

	// TODO: instead of separate array we can use negative numbers in the given to indicate which was processed already
	visited := make([]bool, sz)

	mem := -1
	i := nextNotVisited(visited)
	for {
		if i >= sz {
			break
		}
		if visited[i] {
			i = nextNotVisited(visited)
			continue
		}

		next := index(sz, i)
		//fmt.Printf("i=%d, next=%d\n", i, next)
		if next == i {
			visited[i] = true
			i = nextNotVisited(visited)
			continue
		}

		if mem < 0 {
			mem = nums[next]
			nums[next] = nums[i]
			nums[i] = -1
		} else {
			tmp := nums[next]
			nums[next] = mem
			mem = tmp
		}
		visited[i] = true
		i = next
		//fmt.Printf("mem=%d,  arr=%v, visited=%v\n", mem, nums, visited)
	}
	return nums
}

func index(size int, cur int) int {
	val := (cur * 2) % (size - 1)
	if val == 0 {
		return cur
	}
	return val
}

func nextNotVisited(vis []bool) int {
	for i, v := range vis {
		if !v {
			return i
		}
	}
	return len(vis)
}

func TestShuffleTheArray(t *testing.T) {
	type testCase struct {
		nums     []int
		n        int
		expected []int
	}
	cases := []testCase{
		{[]int{1, 2, 3, 11, 12, 13}, 3, []int{1, 11, 2, 12, 3, 13}},
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 2, 3, 4, 11, 12, 13, 14}, 4, []int{1, 11, 2, 12, 3, 13, 4, 14}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, cs := range cases {
		//fmt.Printf("before: %v\n", cs.nums)
		res := shuffle(cs.nums, cs.n)
		//fmt.Printf("after: %v\n", cs.nums)
		if !equalArrays(res, cs.expected) {
			t.Errorf("shuffle(%v) == %v, expected %v", cs.nums, res, cs.expected)
		}
	}
}

func equalArrays(left []int, right []int) bool {
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
