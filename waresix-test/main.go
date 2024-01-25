package main

import (
	"fmt"
)

func main() {
	s := palindrome("hellosannasmith")
	if len(s) < 3 {
		s = "none"
	}

	token := "ys4lk3zb8"
	i := len(token)
	j := len(s)

	iCounting := 0
	jCounting := 0
	var str string

	for {
		if jCounting < j {
			str += string(s[jCounting])
			jCounting++
		}

		if iCounting < i {
			str += string(token[iCounting])
			iCounting++
		}

		if (i-iCounting == 0) && (j-jCounting == 0) {
			break
		}
	}

	fmt.Println(str)
}

func palindrome(str string) string {
	var start, maxLength = 0, 1
	if len(str) <= 1 {
		return str
	}

	for i := 0; i < len(str)-maxLength/2; i++ {
		var newStart, newLength int
		newStart, newLength = extendPalindrome(str, i, i)
		if newLength > maxLength {
			maxLength = newLength
			start = newStart
		}

		newStart, newLength = extendPalindrome(str, i, i+1)
		if newLength > maxLength {
			maxLength = newLength
			start = newStart
		}
	}

	return str[start : start+maxLength]
}

func extendPalindrome(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return i + 1, j - i - 1
}
