package leetcode

import (
	"testing"
)

func generatePT(numRows int) [][]int {
	res := make([][]int, numRows, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1, i+1)
		for j := 0; j <= i; j++ {
			res[i][j] = pascal(i, j, &res)
		}
		//fmt.Printf("row %v: %v\n", i, res[i])
	}
	return res
}

func pascal(i int, j int, m *[][]int) int {
	//fmt.Printf("\ti=%v, j=%v\n", i, j)
	if i == 0 || j == 0 || i == j {
		return 1
	}
	return (*m)[i-1][j-1] + (*m)[i-1][j]
}

func TestPascalsTriangles(t *testing.T) {
	type testCase struct {
		val      int
		expected [][]int
	}
	cases := []testCase{
		{5, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		},
		},
	}
	for _, cs := range cases {
		res := generatePT(cs.val)
		if !equalM(res, cs.expected) {
			t.Errorf("generatePT(%v) == %v, expected %v", cs.val, res, cs.expected)
		}
	}
}
func equalM(a [][]int, b [][]int) bool {
	alen, blen := len(a), len(b)
	//fmt.Printf("1. alen=%v, blen=%v\n", alen, blen)
	if alen != blen {
		return false
	}
	for i := 0; i < alen; i++ {
		if !equalArr(a[i], b[i]) {
			//fmt.Printf("1.\n\ta: %v\n\tb: %v\n", a, b)
			return false
		}
	}

	return true
}

func equalArr(a []int, b []int) bool {
	alen, blen := len(a), len(b)
	//fmt.Printf("2. alen=%v, blen=%v\n", alen, blen)
	//fmt.Printf("2. a: %v\nb: %v\n", a, b)
	if alen != blen {
		return false
	}
	for i := 0; i < alen; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
