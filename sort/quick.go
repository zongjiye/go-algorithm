package sort

func partition(array []int, begin, end int) int {
	i := begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	j := end       // array[end]是数组的最后一位
	// 没重合之前
	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i]
			j--
		} else {
			i++
		}
	}

	/* 跳出while循环后，i = j。
	 * 此时数组被分割成两个部分  -->  array[begin+1] ~ array[i-1] < array[begin]
	 *                        -->  array[i+1] ~ array[end] > array[begin]
	 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
	 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
	 */
	if array[i] >= array[begin] {
		i--
	}

	array[i], array[begin] = array[begin], array[i]
	return i
}

func QuickSort(array []int, begin, end int) {
	if begin < end {
		loc := partition(array, begin, end)
		QuickSort(array, begin, loc-1)
		QuickSort(array, loc+1, end)
	}
}

// QuickSort1 改进一::小规模插入排序
func QuickSort1(array []int, begin, end int) {
	if end-begin <= 4 {
		insertionSort(array)
		return
	}

	loc := partition(array, begin, end)
	QuickSort(array, begin, loc-1)
	QuickSort(array, loc+1, end)
}

// QuickSort2 三切分的快速排序
func QuickSort2(array []int, begin, end int) {
	if begin < end {
		// 三向切分函数，返回左边和右边下标
		lt, gt := partition3(array, begin, end)
		// 从lt到gt的部分是三切分的中间数列
		// 左边三向快排
		QuickSort2(array, begin, lt-1)
		// 右边三向快排
		QuickSort2(array, gt+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition3(array []int, begin, end int) (int, int) {
	lt := begin       // 左下标从第一位开始
	gt := end         // 右下标是数组的最后一位
	i := begin + 1    // 中间下标，从第二位开始
	v := array[begin] // 基准数

	// 以中间坐标为准
	for i <= gt {
		if array[i] > v { // 大于基准数，那么交换，右指针左移
			array[i], array[gt] = array[gt], array[i]
			gt--
		} else if array[i] < v { // 小于基准数，那么交换，左指针右移
			array[i], array[lt] = array[lt], array[i]
			lt++
			i++
		} else {
			i++
		}
	}

	return lt, gt
}

// QuickSort3 改进三::伪尾递归优化
func QuickSort3(array []int, begin, end int) {
	for begin < end {
		// 进行切分
		loc := partition(array, begin, end)

		// 那边元素少先排哪边
		if loc-begin < end-loc {
			// 先排左边
			QuickSort3(array, begin, loc-1)
			begin = loc + 1
		} else {
			// 先排右边
			QuickSort3(array, loc+1, end)
			end = loc - 1
		}
	}
}
