package models

import "strconv"

// 数位dp
func digitDP(n int) int {
	s := strconv.Itoa(n)
	var fcount func(int, int, bool, bool) int
	fcount = func(i int, mask int, isLimit bool, isNum bool) int {
		if i == len(s) {
			if isNum {
				return 1
			}
			return 0
		}
		res := 0
		if !isNum {
			res = fcount(i+1, mask, false, false)
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for d := 0; d <= up; d++ {
			if mask>>d&1 == 0 {
				res += fcount(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return res
	}
	return fcount(0, 0, true, false)
}
