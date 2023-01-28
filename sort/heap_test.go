package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestHeapSort(t *testing.T) {
	Convey("TestHeapSort", t, func() {
		Convey("TestHeapSort Normal", func() {
			list := []int{10, 88, 2, 199, 1, 199, -19, 99}
			HeapSortN(list)
			log.Println("heap sort normal:", list)
		})

		Convey("TestHeapSort", func() {
			list := []int{10, 88, 2, 199, 1, 199, -19, 99}
			HeapSort(list)
			log.Println("heap sort:", list)
		})
	})
}
