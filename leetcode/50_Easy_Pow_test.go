package leetcode

import (
	"math"
	"testing"
)

func myPow(x float64, a int) float64 {
	if a == 0 {
		return 1.0
	}
	if a < 0 {
		return 1 / myPow(x, -a)
	}
	res := 1.0
	for ; a > 1; a /= 2 {
		if a%2 == 1 {
			res *= x
		}
		x *= x
	}
	return res * x
}

func TestPow(t *testing.T) {
	type testCase struct {
		val      float64
		exponent int
		expected float64
	}
	cases := []testCase{
		{2.0, 2, 4.0},
		{3.0, 3, 27.0},
		{2.0, 10, 1024.0},
		{1, 1000, 1.0},
		{4, -2, 0.0625},
		{2, -2, 0.25},
		{2.10000, 3, 9.26100},
		{2.0000, -2147483648, 0},
	}

	for _, cs := range cases {
		res := myPow(cs.val, cs.exponent)
		if math.Abs(res-cs.expected) >= 0.00000000000001 {
			t.Errorf("myPow(%v, %v) == %v, expected %v", cs.val, cs.exponent, res, cs.expected)
		}
	}
}
