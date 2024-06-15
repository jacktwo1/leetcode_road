package main

func kmp(s1 string, s2 string) int {
	// s1 中当前比对的位置是 i1
	// s2 中当前比对的位置是 i2
	n1, n2 := len(s1), len(s2)
	next := getNext(s2)
	i1, i2 := 0, 0
	for i1 < n1 && i2 < n2 {
		if s1[i1] == s2[i2] {
			i1++
			i2++
		} else if i2 == 0 {
			i1++
		} else {
			i2 = next[i2]
		}
	}

	if i2 == n2 {
		return i1 - n2
	}
	return -1
}

func getNext(s string) []int {
	if len(s) == 1 {
		return []int{-1}
	}
	next := make([]int, len(s)+1)
	next[0], next[1] = -1, 0
	cn := 0 // 表示要和前一个字符比较的下标
	for i := 2; i <= len(s); {
		if s[i-1] == s[cn] {
			next[i] = cn + 1
			i, cn = i+1, cn+1
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}

func main() {

}
