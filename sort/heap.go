package sort

import "algorithm-go/heap"

func HeapSortN(list []int) {
	h := heap.NewHeap(list)

	// 最坏是O(nlog)
	for _, v := range list {
		h.Push(v)
	}
	// 最坏是O(nlog)
	for range list {
		h.Pop()
	}
}

func HeapSort(array []int) {
	// 堆的元素数量
	count := len(array)

	// 最底层的叶子节点下标，该节点位置不定，但是该叶子节点右边的节点都是叶子节点
	start := count/2 + 1

	// 最后的元素下标
	end := count - 1

	// 最底层开始逐一进行下沉
	for start >= 0 {
		sift(array, start, count)
		// 表示左偏移一个节点，如果该层没有节点了，那么表示到了上一层的最右边
		start--
	}

	// 下沉结束 开始排序
	// 元素大于2个的最大堆才可以移除
	for end > 0 {
		// 将堆顶元素与堆尾元素互换，表示移除最大堆元素
		array[end], array[0] = array[0], array[end]
		// 对堆顶进行下沉操作
		sift(array, 0, end)
		// 一直移除堆顶元素
		end--
	}
}

// 下沉操作， push
func sift(array []int, start, count int) {
	// 父亲节点
	root := start

	// 左儿子
	child := 2*root + 1

	// 如果有下一代
	for child < count {
		// 判断一下是否右儿子比左儿子大, 是翻转的变为右儿子
		if count-child > 1 && array[child] < array[child+1] {
			child++
		}

		// 比较父亲节点和孩子节点
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			// 继续下沉
			root = child
			child = 2*root + 1
		} else {
			return
		}

	}
}
