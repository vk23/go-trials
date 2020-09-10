package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

// 1. use strconv package to convert and add two integers
func addBinary1(a string, b string) string {
	aInt, _ := strconv.ParseInt(a, 2, 64)
	bInt, _ := strconv.ParseInt(b, 2, 64)
	sumInt := aInt + bInt
	return strconv.FormatInt(sumInt, 2)
}

// 2. convert manually and add two integers
func addBinary2(a string, b string) string {
	aInt, bInt := 0, 0

	// convert a from string to int
	aLen := len(a)
	for _, r := range a {
		aLen--
		if r == '1' {
			aInt += pow(2, aLen)
		}
	}

	// convert b from string to int
	bLen := len(b)
	for _, r := range b {
		bLen--
		if r == '1' {
			bInt += pow(2, bLen)
		}
	}

	// convert sum back to string
	sumInt := aInt + bInt
	if sumInt == 0 {
		return "0"
	}
	res := ""
	for sumInt > 0 {
		rem := sumInt % 2
		res = fmt.Sprint(rem) + res
		sumInt /= 2
	}
	return res
}

// 3. add binaries without conversion
// TODO: simplify, use bytes and bitwise logic
func addBinary3(a string, b string) string {
	// add leading zeros for convenience
	aLen, bLen := len(a), len(b)
	for i := 0; i < aLen-bLen; i++ {
		b = "0" + b
	}
	for i := 0; i < bLen-aLen; i++ {
		a = "0" + a
	}

	res, x := "", 0
	for i := len(a) - 1; i >= 0; i-- {
		tmp, xtmp := "", 0
		if x > 0 {
			xtmp = 1
			x--
		}
		if a[i] != b[i] {
			tmp = "1"
		} else if a[i] == '0' {
			tmp = "0"
		} else if a[i] == '1' {
			tmp = "0"
			x++
		}

		if xtmp > 0 {
			if tmp == "0" {
				tmp = "1"
			} else {
				tmp = "0"
				x++
			}
		}
		res = tmp + res
	}
	res = fmt.Sprint(x) + res
	// remove leading zeros
	n := -1
	for i, ch := range res {
		if ch == '1' {
			break
		}
		n = i
	}
	res = res[n+1:]
	if res == "" {
		return "0"
	}
	return res
}

func pow(x int, a int) int {
	if a == 0 {
		return 1
	}
	res := 1
	for ; a > 0; a-- {
		res *= x
	}
	return res
}

func TestAddBinary(t *testing.T) {
	type testCase struct {
		a, b     string
		expected string
	}

	var cases = []testCase{
		{"10", "1111", "10001"},
		{"0", "1", "1"},
		{"0", "0", "0"},
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, cs := range cases {
		res := addBinary1(cs.a, cs.b)
		if res != cs.expected {
			t.Errorf("addBinary1(%v, %v) == %v, expected %v", cs.a, cs.b, res, cs.expected)
		}
		res = addBinary2(cs.a, cs.b)
		if res != cs.expected {
			t.Errorf("addBinary2(%v, %v) == %v, expected %v", cs.a, cs.b, res, cs.expected)
		}
		res = addBinary3(cs.a, cs.b)
		if res != cs.expected {
			t.Errorf("addBinary3(%v, %v) == %v, expected %v", cs.a, cs.b, res, cs.expected)
		}
	}
}
