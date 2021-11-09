package second_day

import (
	"fmt"
)


/*iven a string containing digits from 2-9 inclusive, return all possible
letter combinations that the number could represent. Return the answer in any order.*/

func letterCombinations(digits string) []string {
	m := map[string][]string{}
	m["2"] = []string{"a", "b", "c"}
	m["3"] = []string{"d", "e", "f"}
	m["4"] = []string{"g", "h", "i"}
	m["5"] = []string{"j", "k", "l"}
	m["6"] = []string{"m", "n", "o"}
	m["7"] = []string{"p", "q", "r", "s"}
	m["8"] = []string{"t", "u", "v"}
	m["9"] = []string{"w", "x", "y", "z"}
	a_digits := []string{}
	a_values := []string{}
	result := []string{}
	for _, i := range digits {
		s := string(i)
		a_digits = append(a_digits,s)
		a_values = append(a_values,m[s]...)
	}

	for n, i := range a_values{
		for _, j := range a_values[1+n:]{
			result = append(result, i+j)
		}
	}
	return result
}

func test_1()  {
	s := "23"
	letterCombinations(s)
	fmt.Println(letterCombinations(s))
}


/*Given a string s, return the longest palindromic substring in s.*/

func longestPalindrome(s string) string {
	maxLen := len(s)

	for l := maxLen; l > 0; l-- {
		for i := 0; i + l <= maxLen; i++ {
			lp, rp := i, i + l -1

			for lp < rp && s[lp] == s[rp] {
				lp++
				rp--
			}

			if lp >= rp {
				return s[i:i+l]
			}
		}
	}

	return s[0:1]
}

func test_2() {
	fmt.Println(longestPalindrome("sbba"))
}