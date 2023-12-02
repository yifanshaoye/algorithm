package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(constructArray2(10,4))
}
func flipLights(n int, presses int) int {
	seen := make(map[int]struct{})
	for i:=0; i<8; i++ {
		parr := [3]int{}
		sum := 0
		for j:=0; j<3; j++ {
			parr[j] = i>>j & 1
			sum += parr[j]
		}
		if sum%2 == presses%2 && sum <= presses {
			status := parr[0] ^ parr[2] ^ parr[3]
			if n >= 2 {
				status |= (parr[0] ^ parr[1]) << 1
			}
			if n >= 3 {
				status |= (parr[0] ^ parr[2]) << 2
			}
			seen[status] = struct{}{}
		}
	}
	return len(seen)
}


func checkValidString(s string) bool {
	min, max := 0, 0
	for i:=0; i<len(s); i++ {
		if s[i]=='(' {
			min++
			max++
		} else if s[i] == ')' {
			if min>0 {
				min--
			}
			max--
			if max<0 {
				return false
			}
		} else {
			if min>0 {
				min--
			}
			max++
		}
	}
	return min==0
}