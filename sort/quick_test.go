package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestQuickSort(t *testing.T) {
	Convey("TestQuickSort", t, func() {
		Convey("TestQuickSort Normal", func() {
			list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
			QuickSort(list, 0, len(list)-1)
			log.Println(list)
		})

		Convey("TestQuickSort ", func() {
			list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
			QuickSort(list, 0, len(list)-1)
			log.Println(list)
		})

		Convey("TestQuickSort Normal", func() {
			list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
			QuickSort(list, 0, len(list)-1)
			log.Println(list)
		})

		Convey("TestQuickSort Normal", func() {
			list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
			QuickSort(list, 0, len(list)-1)
			log.Println(list)
		})
	})
}
