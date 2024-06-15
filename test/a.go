package test

func SingleNumber(nums []int) [2]int {

	x := 0
	for _, num := range nums {
		x ^= num
	}
	fristOne := x & -x
	a, b := 0, 0
	for _, num := range nums {
		if num&fristOne == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return [2]int{a, b}
}
