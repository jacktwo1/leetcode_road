package main

import "fmt"

//type KthLargest struct {
//	data []int
//	size int
//}
//
//func heapInsert(arr []int, index int) {
//	for arr[index] < arr[(index-1)/2] {
//		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
//		index = (index - 1) / 2
//	}
//}
//
//func heapify(arr []int, index int, heapSize int) {
//	left := index*2 + 1
//	for left < heapSize {
//		smallest := left
//		if left+1 < heapSize && arr[left+1] < arr[left] {
//			smallest = left + 1
//		}
//		if arr[smallest] > arr[index] {
//			smallest = index
//		}
//		if smallest == index {
//			break
//		}
//		arr[index], arr[smallest] = arr[smallest], arr[index]
//		index = smallest
//		left = 2*index + 1
//	}
//}
//
//func Constructor(k int, nums []int) KthLargest {
//	heap := KthLargest{}
//	heap.size = k
//	for i := 0; i < len(nums); i++ {
//		heap.Add(nums[i])
//	}
//
//	return heap
//}
//
//func (this *KthLargest) Add(val int) int {
//	if len(this.data) < this.size-1 {
//		this.data = append(this.data, val)
//	} else if len(this.data) == this.size-1 {
//		this.data = append(this.data, val)
//		for i := (this.size - 2) / 2; i >= 0; i-- {
//			heapify(this.data, i, this.size)
//		}
//	} else {
//		if val > this.data[0] {
//			this.data[0] = val
//			heapify(this.data, 0, this.size)
//		}
//	}
//	return this.data[0]
//
//}

func topKFrequent(nums []int, k int) []int {

	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	h := new(heap)
	for k, v := range numMap {
		h.data = append(h.data, [2]int{k, v})
		h.adjustUp(h.size)
		h.size++
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = h.pop()
	}
	return res
}

type heap struct {
	data [][2]int
	size int
}

func (h *heap) adjustUp(index int) {
	for h.data[index][1] > h.data[(index-1)/2][1] {
		h.data[index], h.data[(index-1)/2] = h.data[(index-1)/2], h.data[index]
		index = (index - 1) / 2
	}
}

func (h *heap) adjustDown(index int, size int) {
	left := 2*index + 1
	for left < size {
		largerest := left
		if left+1 < size && h.data[left+1][1] > h.data[left][1] {
			largerest = left + 1
		}

		if h.data[index][1] > h.data[largerest][1] {
			largerest = index
		}

		if largerest == index {
			return
		}

		h.data[index], h.data[largerest] = h.data[largerest], h.data[index]
		index = largerest
		left = 2*index + 1
	}
}

func (h *heap) pop() int {
	h.data[0], h.data[h.size-1] = h.data[h.size-1], h.data[0]
	h.adjustDown(0, h.size-1)
	h.size--
	return h.data[h.size-1][0]
}
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	a := topKFrequent(nums, k)
	fmt.Println(a)
}
