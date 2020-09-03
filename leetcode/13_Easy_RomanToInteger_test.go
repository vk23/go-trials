package leetcode

import (
	"testing"
)

var romans = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(str string) int {
	runes := []rune(str)
	res := 0

	var prev, curr int
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		curr = romans[string(r)]
		if curr < prev {
			res -= curr
		} else {
			res += curr
		}
		prev = curr
	}
	return res
}

func TestRomanToInt(t *testing.T) {
	type testCase struct {
		value    string
		expected int
	}
	cases := []testCase{
		{"V", 5},
		{"I", 1},
		{"IV", 4},
		{"IX", 9},
		{"XXIII", 23},
		{"LVII", 57},
		{"MCMXCIX", 1999},
	}

	for _, cs := range cases {
		res := romanToInt(cs.value)
		if res != cs.expected {
			t.Errorf("RomanToIn(%v) == %v, expected %v", cs.value, res, cs.expected)
		}
	}
}
