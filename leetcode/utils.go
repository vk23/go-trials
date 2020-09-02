package leetcode

import "errors"

var NotImplementedError = errors.New("not implemented")

func Pow(x int, a int) (int, error) {
	if a < 0 {
		return 0, NotImplementedError
	}
	if a == 0 {
		return 1, nil
	}
	res := 1
	for ; a > 0; a-- {
		res *= x
	}
	return res, nil
}
