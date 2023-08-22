package main

import (
	"fmt"
)

type Heap struct {
	array []int
	size  int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Init(array []int) {
	n := len(array)
	h.array = array
	h.size = n
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *Heap) Len() int {
	return h.size
}

func (h *Heap) Less(i, j int) bool {
	return h.array[i] > h.array[j]
}

func (h *Heap) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *Heap) Pop() int {
	n := h.Len() - 1
	// swap 根结点和最后一个节点
	h.Swap(0, n)
	// down 根节点下沉
	h.down(0, n)
	// 弹出最后一个元素
	x := h.array[n]
	h.array = h.array[:n]
	h.size--
	return x
}

func (h *Heap) down(i0, n int) {
	i := i0
	for {
		j := 2*i + 1
		// 左孩子不存在, j1 < 0 after int overflow
		if j >= n || j < 0 {
			break
		}
		// 右孩子存在, 右孩子 > 左孩子
		if j+1 < n && h.Less(j+1, j) {
			j++
		}
		// 判断当前节点是否大于左、右孩子最大值, 如果大于停止下沉
		if h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

func (h *Heap) Push(x int) {
	h.array = append(h.array, x)
	h.size++
	// 上浮
	h.up(h.Len() - 1)
}

func (h *Heap) up(i int) {
	for {
		parent := (i - 1) / 2
		// 如果 i == parent(没有父节点)， h.array[parent] > h.array[i] 退出上浮
		if i == parent || h.Less(parent, i) {
			break
		}
		// swap
		h.Swap(i, parent)
		i = parent
	}
}

func main() {
	a := []int{2, 5, 9, 3, 1}
	h := NewHeap()
	h.Init(a)
	h.Push(20)
	fmt.Println("max num:", h.array[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	fmt.Println("\nover")
}
