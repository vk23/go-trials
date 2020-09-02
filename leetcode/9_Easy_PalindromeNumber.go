package main

import (
	"fmt"
	"math"
)

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
Coud you solve it without converting the integer to a string?
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// find the leftmost digit in number
	div := 1000000000
	for ; div > 1 && x/div < 1; div = div / 10 {
	}

	length := int(math.Log10(float64(div)))
	//fmt.Printf("div=%d, len=%d\n", div, length)
	if length == 0 {
		return true
	}

	for i := length; i > 0 && x > 0; i = i - 2 {
		d := pow(10, i)
		l := x / d
		r := x % 10

		//fmt.Printf("x=%d,d=%d [%d;%d]\n", x, d, l, r)
		if l != r {
			return false
		}
		x = x % d
		x = x / 10
	}

	return true
}

func pow(x int, a int) int {
	if a < 0 {
		return -1 // not implemented
	}
	if a == 0 {
		return 1
	}
	res := 1
	for ; a > 0; a-- {
		res *= x
	}
	return res
}

func main() {
	for _, cs := range cases {
		fmt.Printf("Case %10v, passed = %v\n", cs, cs.res == isPalindrome(cs.num))
	}
}

type Case struct {
	num int
	res bool
}

var cases = []Case{
	{121, true},
	{11, true},
	{10, false},
	{-121, false},
	{65456, true},
	{65476, false},
	{112211, true},
}
