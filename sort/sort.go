package sort

// BubbleSort 冒泡排序 稳定 交换类排序算法
func BubbleSort(list []int) {
	n := len(list)
	// 在一轮中有没有交换过
	isSwap := false
	// 进行 N-1 轮迭代
	for i := n - 1; i > 0; i-- {
		// 每次从第一位开始比较，比较到第 i 位就不比较了，因为前一轮该位已经有序了
		for j := 0; j < i; j++ {
			// 如果前面的数比后面的大，那么交换
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				isSwap = true
			}
		}
		// 如果在一轮中没有交换过，那么已经排好序了，直接返回
		if !isSwap {
			return
		}
	}
}

// SelectSort 选择排序 不稳定 选择类排序算法
func SelectSort(list []int) {
	n := len(list)
	for i := 0; i < n-1; i++ {
		// 每次从第 i 位开始，找到最小的元素
		min := list[i]
		minIndex := i
		for j := i + 1; j < n; j++ {
			// 如果找到的数比上次的还小，那么最小的数变为它
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}

		if i != minIndex {
			// 这一轮找到的最小数的下标不等于最开始的下标，交换元素
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

// SelectGoodSort 选择排序优化 不稳定 选择类排序算法
func SelectGoodSort(list []int) {
	n := len(list)
	// 只需循环一半
	for i := 0; i < n/2; i++ {
		minIndex := i
		maxIndex := i
		for j := i + 1; j < n-i; j++ {
			// 找到最大值下标
			if list[j] > list[maxIndex] {
				maxIndex = j
				continue
			}
			// 找到最小值下标
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		// 如果最大值是开头的元素，而最小值不是最尾的元素
		if maxIndex == i && minIndex != n-i-1 {
			// 先将最大值和最尾的元素交换
			list[maxIndex], list[n-i-1] = list[n-i-1], list[maxIndex]
			// 然后最小的元素放在最开头
			list[minIndex], list[i] = list[i], list[minIndex]
		} else if maxIndex == i && minIndex == n-i-1 {
			// 如果最大值在开头，最小值在结尾，直接交换
			list[maxIndex], list[minIndex] = list[minIndex], list[maxIndex]
		} else {
			// 否则先将最小值放在开头，再将最大值放在结尾
			list[minIndex], list[i] = list[i], list[minIndex]
			list[maxIndex], list[n-i-1] = list[n-i-1], list[maxIndex]
		}
	}
}

// InsertSort 直接插入排序 稳定 插入类排序算法
func InsertSort(list []int) {
	n := len(list)

	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if deal < list[j] {
			for ; j >= 0 && deal < list[j]; j-- {
				// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
				list[j+1] = list[j]
			}
			// 结束时j指向空位前一位, 待排序的数插入空位（j+1）
			list[j+1] = deal
		}
	}
}

// ShellSort 希尔排序 不稳定 插入类排序算法
func ShellSort(list []int) {
	n := len(list)

	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		// i 待排序的元素
		for i := step; i < n; i += step {
			// j 在比较过程中, 待排序元素的位置
			for j := i; j >= step; j -= step {
				// 同组左边的元素 <= 待排序元素
				if list[j-step] <= list[j] {
					break
				}
				// 交换
				list[j-step], list[j] = list[j], list[j-step]
			}
		}
	}
}
