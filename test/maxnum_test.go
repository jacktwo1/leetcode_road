package test

import "testing"

func TestMaxNum(t *testing.T) {
	a := []int{3, 6, 8}
	limit := 65632
	if res := MaxNum(limit, a); res != 6666 {
		t.Errorf("单元测试失败, 测试输入为 %v,测试输出为 %d", a, res)
	}
}
