package main

import "fmt"

func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

func heapify(arr []int, index int, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		smallest := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			smallest = left + 1
		}
		if arr[smallest] < arr[index] {
			smallest = index
		}
		if smallest == index {
			break
		}
		arr[index], arr[smallest] = arr[smallest], arr[index]
		index = smallest
		left = 2*index + 1
	}
}

func heapSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		heapInsert(nums, i)
	}

	hl := len(nums)
	hl -= 1
	nums[0], nums[hl] = nums[hl], nums[0]

	for hl > 0 {
		heapify(nums, 0, hl)
		hl -= 1
		nums[0], nums[hl] = nums[hl], nums[0]
	}

}

func main() {
	a := []int{1, 6, 7, 4, 3}
	heapSort(a)
	fmt.Println(a)
}
