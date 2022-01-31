package main

func main() {
}

func isValid(s string) bool {
	st := []rune{}
	bracketsMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range s {
		if len(st) == 0 {
			st = append(st, v)
			continue
		}
		if bracketsMap[v] == st[len(st)-1] {
			st = st[:len(st)-1]
		} else {
			st = append(st, v)
		}
	}
	return len(st) == 0
}
