package heap

// 优先队列是一种能完成以下任务的队列：插入一个数值，取出最小或最大的数值（获取数值，并且删除）。
// 优先队列可以用二叉树来实现，我们称这种结构为二叉堆。
// 最小堆和最大堆是二叉堆的一种，是一棵完全二叉树（一种平衡树）。

// 最小堆 :: 父节点的值都小于左右儿子节点。
// 最大堆 :: 父节点的值都大于左右儿子节点。

// Heap 最大堆的实现
// 一个最大堆，一棵完全二叉树
// 最大堆有两个核心操作，一个是上浮，一个是下沉，分别对应 push 和 pop
type Heap struct {
	Size int
	// 使用内部的数组来模拟树
	// 一个节点下标为 i，那么父亲节点的下标为 (i-1)/2
	// 一个节点下标为 i，那么左儿子的下标为 2i+1，右儿子下标为 2i+2
	Array []int
}

func NewHeap() *Heap {
	return &Heap{Array: make([]int, 0)}
}

// Push 上浮操作
func (h *Heap) Push(x int) {
	// push
	h.Array = append(h.Array, x)
	h.Size++
	// up
	h.up(h.Len() - 1)
}

// Pop 下沉操作
func (h *Heap) Pop() int {
	n := h.Len() - 1
	h.Swap(0, n)
	h.down(0, n)

	// pop
	temp := h.Array[n]
	h.Array = h.Array[:n]
	h.Size--
	return temp
}

func (h *Heap) Len() int {
	return h.Size
}

func (h *Heap) up(i int) {
	for {
		parent := (i - 1) / 2 // parent
		if i == parent || !h.Less(i, parent) {
			break
		}
		h.Swap(i, parent)
		i = parent
	}
}

func (h *Heap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // left
		// find max index in left of right child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2
		}
		// when left and child less than parent
		if !h.Less(j, i) {
			break
		}
		// swap
		h.Swap(i, j)
		i = j
	}

	return i > i0
}

// Less 最大堆的实现
func (h *Heap) Less(i, j int) bool {
	return h.Array[i] < h.Array[j]
}

// Swap 交换
func (h *Heap) Swap(i, j int) {
	h.Array[i], h.Array[j] = h.Array[j], h.Array[i]
}
