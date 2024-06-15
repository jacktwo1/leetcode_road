package test

import "testing"

func TestSingleNumber(t *testing.T) {
	a := []int{1, 2, 1, 4, 2, 5}
	if res := SingleNumber(a); res != [2]int{5, 3} && res != [2]int{3, 5} {
		t.Errorf("单元测试失败, 测试输入为%v, 测试输出为%v", a, res)
	}
}
