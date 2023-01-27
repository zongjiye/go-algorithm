package sort

// mergex return Slice
func mergex(left []int, right []int) []int {
	res := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}

// MergeSortTopX return slice
func MergeSortTopX(list []int) []int {
	n := len(list)
	if n <= 1 {
		return list
	}
	left := MergeSortTopX(list[:n/2])
	right := MergeSortTopX(list[n/2:])
	return mergex(left, right)
}

// merge not return
func merge(array []int, begin, mid, end int) {
	leftSize := mid - begin
	rightSize := end - mid
	newSize := leftSize + rightSize
	res := make([]int, 0, newSize)

	l, r := 0, 0
	for l < leftSize && r < rightSize {
		if array[begin+l] < array[mid+r] {
			res = append(res, array[begin+l])
			l++
		} else {
			res = append(res, array[mid+r])
			r++
		}
	}

	res = append(res, array[begin+l:mid]...)
	res = append(res, array[mid+r:end]...)

	for i := 0; i < newSize; i++ {
		array[begin+i] = res[i]
	}
}

// MergeSortTop top-down
func MergeSortTop(array []int, begin, end int) {
	// 元素数量大于1时才进入递归
	if end-begin > 1 {
		// 将数组一分为二，分为 array[begin,mid) 和 array[mid,end)
		mid := begin + (end-begin+1)/2
		// 左边排序
		MergeSortTop(array, begin, mid)
		// 右边排序
		MergeSortTop(array, mid, end)
		// 两个有序数组进行合并'
		merge(array, begin, mid, end)
	}
}

// MergeSortBottom Bottom-up
func MergeSortBottom(array []int, begin, end int) {
	// 步数为1开始，step长度的数组表示一个有序的数组
	step := 1
	// 范围大于 step 的数组才可以进入归并
	for end-begin > step {
		// 从头到尾对数组进行归并操作
		// step << 1 = 2 * step 表示偏移到后两个有序数组将它们进行归并
		for i := begin; i < end; i += step << 1 {
			var lo = i                // 第一个有序数组的上界
			var mid = lo + step       // 第一个有序数组的下界，第二个有序数组的上界
			var hi = lo + (step << 1) // 第二个有序数组的下界
			// 不存在第二个数组，直接返回
			if mid > end {
				return
			}
			// 第二个数组长度不够
			if hi > end {
				hi = end
			}
			// 两个有序数组进行合并
			merge(array, lo, mid, hi)
		}
		// 上面的 step 长度的两个数组都归并成一个数组了，现在步长翻倍
		step <<= 1
	}

}

// 手摇算法，将 array[l,l+1,l+2,...,mid-2,mid-1,mid,mid+1,mid+2,...,r-2,r-1,r] 从mid开始两边交换位置
// 1.先逆序前部分：array[mid-1,mid-2,...,l+2,l+1,l]
// 2.后逆序后部分：array[r,r-1,r-2,...,mid+2,mid+1,mid]
// 3.上两步完成后：array[mid-1,mid-2,...,l+2,l+1,l,r,r-1,r-2,...,mid+2,mid+1,mid]
// 4.整体逆序： array[mid,mid+1,mid+2,...,r-2,r-1,r,l,l+1,l+2,...,mid-2,mid-1]
func rotation(array []int, l, mid, r int) {
	reverse(array, l, mid-1)
	reverse(array, mid, r)
	reverse(array, l, r)
}

func reverse(array []int, l, r int) {
	for l < r {
		// 左右互相交换
		array[l], array[r] = array[r], array[l]
		l++
		r--
	}
}

// insertionSort
func insertionSort(list []int) {
	n := len(list)
	// 进行 N-1 轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if deal < list[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j] // 某数后移，给待排序留空位
			}
			list[j+1] = deal // 结束了，待排序的数插入空位
		}
	}
}

// mergeBest
func mergeBest(array []int, begin, mid, end int) {
	// 三个下标，将数组 array[begin,mid] 和 array[mid,end-1]进行原地归并
	i, j, k := begin, mid, end-1 // 因为数组下标从0开始，所以 k = end-1

	for j-i > 0 && k-j >= 0 {
		step := 0
		// 从 i 向右移动，找到第一个 array[i]>array[j]的索引
		for j-i > 0 && array[i] <= array[j] {
			i++
		}

		// 从 j 向右移动，找到第一个 array[j]>array[i]的索引
		for k-j >= 0 && array[j] <= array[i] {
			j++
			step++
		}

		// 进行手摇翻转，将 array[i,mid] 和 [mid,j-1] 进行位置互换
		// mid 是从 j 开始向右出发的，所以 mid = j-step
		rotation(array, i, j-step, j-1)
		i = i + step
	}
}

// MergeSortBest 自底向上归并排序优化版本
func MergeSortBest(array []int, n int) {
	// 按照三个元素为一组进行小数组排序，使用直接插入排序
	blockSize := 3
	a, b := 0, blockSize
	for b <= n {
		insertionSort(array[a:b])
		a = b
		b += blockSize
	}
	insertionSort(array[a:n])

	// 将这些小数组进行归并
	for blockSize < n {
		a, b = 0, 2*blockSize
		for b <= n {
			mergeBest(array, a, a+blockSize, b)
			a = b
			b += 2 * blockSize
		}
		if m := a + blockSize; m < n {
			mergeBest(array, a, m, n)
		}
		blockSize *= 2
	}
}
