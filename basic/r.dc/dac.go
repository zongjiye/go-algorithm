package r_dc

// BinarySearch 二分查找
func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		return -1
	}

	// mid = (l + r) / 2
	mid := (l + r) / 2
	midNum := array[mid]

	if midNum == target {
		return 1
	} else if midNum > target {
		// target smaller than mid
		// r = mid + 1
		return BinarySearch(array, target, l, mid-1)
	} else {
		// target bigger than mid
		// l = mid + 1
		return BinarySearch(array, target, mid+1, r)
	}
}

// BinarySearch2 非递归写法
func BinarySearch2(array []int, target int, l, r int) int {
	lTemp := l
	rTemp := r
	for {
		if lTemp > rTemp {
			return -1
		}

		mid := lTemp + rTemp
		midNum := array[mid]
		if midNum == target {
			return 1
		} else if target > midNum {
			lTemp = mid + 1
		} else {
			rTemp = mid - 1
		}
	}
}
