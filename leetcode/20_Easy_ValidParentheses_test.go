package leetcode

import (
	"errors"
	"testing"
)

func isValidParentheses(s string) bool {
	openPar := map[rune]int{
		'(': 1,
		'{': 2,
		'[': 3,
	}
	closePar := map[rune]int{
		')': 1,
		'}': 2,
		']': 3,
	}
	var stack Stack
	for _, r := range s {
		if _, ok := openPar[r]; ok {
			stack.Push(r)
		} else if i, ok := closePar[r]; ok {
			pop, err := stack.Pop()
			if err != nil || openPar[pop] != i {
				return false
			}
		}
	}
	return stack.size == 0
}

func TestValidParentheses(t *testing.T) {
	type testCase struct {
		value    string
		expected bool
	}
	cases := []testCase{
		{"()", true},
		{"()[]{}", true},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, cs := range cases {
		res := isValidParentheses(cs.value)
		if res != cs.expected {
			t.Errorf("ValidParentheses(%v) == %v, expected %v", cs.value, res, cs.expected)
		}
	}
}

type Stack struct {
	runes []rune
	size  int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Peek() (rune, error) {
	if s.size == 0 {
		return -1, errors.New("empty stack")
	}
	return s.runes[s.size-1], nil
}

func (s *Stack) Push(r rune) {
	s.runes = append(s.runes, r)
	s.size++
	//fmt.Printf("size=%d\n", s.size)
}

func (s *Stack) Pop() (rune, error) {
	if s.size == 0 {
		return -1, errors.New("empty stack")
	}
	r := s.runes[s.size-1]
	s.size--
	//fmt.Printf("size=%d\n", s.size)
	s.runes = s.runes[0:s.size]
	return r, nil
}
