package models

// 双端队列 752 open the lock
func openLock(deadends []string, target string) int {
	dmap := make(map[string]bool)
	for _, st := range deadends {
		dmap[st] = true
	}
	if dmap["0000"] {
		return -1
	}

	que1 := []string{"0000"}
	visited1 := map[string]bool{"0000": true}
	que2 := []string{target}
	visited2 := map[string]bool{target: true}
	step := -1
	for len(que1) > 0 && len(que2) > 0 {
		step++
		if len(que1) <= len(que2) {
			qlen := len(que1)
			for q := 0; q < qlen; q++ {
				state := que1[q]
				if visited2[state] {
					return step
				}
				codes := getNextcode(state)
				for _, code := range codes {
					if !visited1[code] && !dmap[code] {
						visited1[code] = true
						que1 = append(que1, code)
					}
				}
			}
			que1 = que1[qlen:]
		} else {
			qlen := len(que2)
			for q := 0; q < qlen; q++ {
				state := que2[q]
				if visited1[state] {
					return step
				}
				codes := getNextcode(state)
				for _, code := range codes {
					if !visited2[code] && !dmap[code] {
						visited2[code] = true
						que2 = append(que2, code)
					}
				}
			}
			que2 = que2[qlen:]
		}
	}
	return -1
}

func getNextcode(state string) []string {
	ret := []string{}
	s := []byte(state)
	for i, b := range s {
		s[i] = b - 1
		if s[i] < '0' {
			s[i] = '9'
		}
		ret = append(ret, string(s))
		s[i] = b + 1
		if s[i] > '9' {
			s[i] = '0'
		}
		ret = append(ret, string(s))
		s[i] = b
	}

	return ret
}
