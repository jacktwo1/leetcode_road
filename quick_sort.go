package main

import (
	"fmt"
	"math/rand"
)

func partition(nums []int, L int, R int) [2]int {
	less := L - 1
	more := R // 这个是大于区间的右边界，因为 nums[R] 作为 partition 位置，只要考虑 [L:R-1] 的位置就好
	index := L
	for index < more {
		if nums[index] < nums[R] {
			nums[index], nums[less+1] = nums[less+1], nums[index]
			index++
			less++
		} else if nums[index] > nums[R] {
			nums[index], nums[more-1] = nums[more-1], nums[index]
			more--
		} else {
			index++
		}
	}
	nums[more], nums[R] = nums[R], nums[more]
	return [2]int{less + 1, more}
}

func quickSort(nums []int, L int, R int) {
	if L < R {
		// 随机选取一个元素和最右边的交换，以最右边的元素partition
		randSelect := L + rand.Intn(R-L)
		nums[randSelect], nums[R] = nums[R], nums[randSelect]
		border := partition(nums, L, R)
		quickSort(nums, L, border[0]-1)
		quickSort(nums, border[1]+1, R)
	}
}

func main() {
	a := []int{1, 23, 46, 44, 6, 73, 54, 75, 56, 3, 34, 657, 8, 86, 4555, 7, 5, 45, 72, 34, 67, 7}
	partition(a, 0, len(a)-1)
	fmt.Println(a)
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
