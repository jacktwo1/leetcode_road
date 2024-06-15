package main

import "testing"

func singleNumber(nums []int) [2]int {

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

func TestSingleNumber(t *testing.T) {
	a := []int{1, 2, 1, 3, 2, 5}

	if res := singleNumber(a); res != [2]int{5, 3} && res != [2]int{3, 5} {
		t.Errorf("单元测试失败, 测试输入为%v, 测试输出为 %v", a, res)
	}
}
