package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestMergeSort(t *testing.T) {
	Convey("TestMergeSort", t, func() {
		Convey("TestMergeSortTopX", func() {
			list := []int{10, 88, 2, 199, 1, 199, -19, 99}
			res := MergeSortTopX(list)
			log.Println(res)
		})

		Convey("TestMergeSortTop", func() {
			list := []int{10, 88, 2, 199, 1, 199, -19, 99}
			MergeSortTop(list, 0, len(list))
			log.Println(list)
		})

		Convey("TestMergeSortBottom", func() {
			list := []int{10, 88, 2, 199, 1, 199, -19, 99}
			MergeSortBottom(list, 0, len(list))
			log.Println(list)
		})

		Convey("TestMergeSortBest", func() {
			list := []int{10, 88, 2, 199, 1, 199, -19, 99}
			MergeSortBest(list, len(list))
			log.Println(list)
		})
	})
}
