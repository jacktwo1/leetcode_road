package test

import (
	"sort"
)

func MaxNum(limit int, nums []int) int {
	sort.Ints(nums)
	limit-- // 这样就是取 <= limit 最大的数

	offset := 1
	for offset <= limit/10 {
		offset *= 10
	}
	res := process(limit, offset, nums)
	if res != -1 {
		return res
	} else {
		return rest(nums, offset/10)
	}

}

func process(limit, offset int, nums []int) int {
	if offset == 0 { // 说明 nums 可以取到等于 limit 的数
		return limit
	}

	cur := (limit / offset) % 10
	nearNum := near(nums, cur)
	if nearNum == -1 {
		return -1
	} else if nums[nearNum] == cur { // 找出来的数字，真的和当前数字 cur 一样
		res := process(limit, offset/10, nums)
		if res != -1 {
			return res
		} else if nearNum > 0 { // 后续没成功，但是自己有下降空间
			nearNum--
			return (limit/(offset*10))*offset*10 + (nums[nearNum] * offset) + rest(nums, offset/10)
		} else { // 后续没成功，自己也没有下降的空间
			return -1
		}
	} else {
		return (limit/(offset*10))*offset*10 + (nums[nearNum] * offset) + rest(nums, offset/10)
	}
}

func near(nums []int, target int) int {
	res := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			res = mid
			left = mid + 1
		}
	}
	return res
}

func rest(nums []int, offset int) int {
	res := 0
	for offset > 0 {
		res += offset * nums[len(nums)-1]
		offset /= 10
	}
	return res
}
