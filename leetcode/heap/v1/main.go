package main

import (
	"fmt"
)

type Heap struct {
	size  int
	array []int
}

func NewHeap(array []int) *Heap {
	return &Heap{array: array}
}

func (h *Heap) Push(x int) {
	// size = 0 直接赋值
	if h.size == 0 {
		h.array[0] = x
		h.size++
		return
	}

	// 记录要插入元素的位置
	i := h.size
	for i > 0 {
		// 父节点 parent
		parent := (i - 1) / 2
		// 父节点大于等于x结束上浮
		if h.array[parent] >= x {
			break
		}
		// swap
		h.array[i] = h.array[parent]
		i = parent
	}
	// 上浮结束,  将x插入到对应位置
	h.array[i] = x
	h.size++
}

func (h *Heap) Pop() int {
	if h.size == 0 {
		return -1
	}

	// 取出根节点
	head := h.array[0]
	h.size--
	x := h.array[h.size]
	h.array[h.size] = head
	// 将最后一个位置的值取出, 进行下沉
	i := 0
	for {
		// 左孩子的下标
		j := 2*i + 1
		// 没有左孩子
		if j >= h.size {
			break
		}
		//  存在右孩子且右孩子 > 左左孩子，记录右孩子的下标
		if j+1 < h.size && h.array[j+1] > h.array[j] {
			j++
		}
		// x大于等于左、右孩子的最大值，下沉结束
		if x >= h.array[j] {
			break
		}
		// swap
		h.array[i] = h.array[j]
		i = j
	}
	// 停止下沉, 当前i的位置为最后元素的的位置
	h.array[i] = x

	return head
}

// 堆排序
func heapSort(array []int) {
	h := NewHeap(array)
	// 进堆
	for _, v := range array {
		h.Push(v)
	}
	// 出堆
	for range array {
		h.Pop()
	}
}

func main() {
	a := []int{2, 5, 9, 3, 1}
	heapSort(a)
	fmt.Println(a)
}
