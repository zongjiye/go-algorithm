package r_dc

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	Convey("TestBinarySearch-Result", t, func() {
		array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
		right := len(array) - 1
		Convey("BinarySearch should be return 1", func() {
			So(BinarySearch(array, 89, 0, right), ShouldEqual, 1)
		})

		Convey("BinarySearch should be return -1", func() {
			So(BinarySearch(array, 99, 0, right), ShouldEqual, -1)
		})

		Convey("BinarySearch2 should be return 1", func() {
			So(BinarySearch2(array, 189, 0, right), ShouldEqual, 1)
		})

		Convey("BinarySearch2 should be return -1", func() {
			So(BinarySearch2(array, 223, 0, right), ShouldEqual, -1)
		})

	})
}
