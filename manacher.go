package main

func countSubstrings(s string) int {
	n := len(s)
	expandString := make([]byte, 2*n+1)
	index := 0
	for i := 0; i <= 2*n; i++ {
		if i&1 == 0 {
			expandString[i] = '*'
		} else {
			expandString[i] = s[index]
			index++
		}
	}

	count := make([]int, 2*n+1)
	res := 0
	maxR, c := 0, 0
	for i := 0; i <= 2*n; i++ {
		length := 1
		if maxR > i {
			length = min(count[2*c-i], maxR-i)
		}

		for i+length <= 2*n && i-length >= 0 && expandString[i+length] == expandString[i-length] {
			length++
		}

		if i+length > maxR {
			maxR = i + length
			c = i
		}
		count[i] = length
		res += length / 2
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
