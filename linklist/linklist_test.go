package linklist

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestLinkList(t *testing.T) {
	Convey("TestLinkList", t, func() {
		head := NewLinkList()
		head.TailInsertion([]int{1, 2, 3})

		Convey("TestHeadInsertion", func() {
			head.HeadInsertion([]int{1, 2, 3})
			So(head.Travel()[0], ShouldEqual, 3)
			So(head.Travel()[1], ShouldEqual, 2)
			So(head.Travel()[2], ShouldEqual, 1)
		})

		Convey("TestTailInsertion", func() {
			head.TailInsertion([]int{1, 2, 3})
			So(head.Travel()[0], ShouldEqual, 1)
			So(head.Travel()[1], ShouldEqual, 2)
			So(head.Travel()[2], ShouldEqual, 3)
		})

		Convey("TestLength", func() {
			So(head.Length(), ShouldEqual, 3)
		})

		Convey("TestTravel", func() {
			So(head.Travel()[0], ShouldEqual, 1)
			So(head.Travel()[1], ShouldEqual, 2)
			So(head.Travel()[2], ShouldEqual, 3)
		})

		Convey("TestInsert", func() {
			node := new(LinkNode)
			node.Data = 4
			head.Insert(1, node)
			log.Println("after insert:", head.Travel())
			So(head.Travel()[0], ShouldEqual, 4)
		})

		Convey("TestDelete", func() {
			head.Delete(1)
			log.Println("after delete:", head.Travel())
			So(head.Travel()[0], ShouldEqual, 2)
		})
		Convey("TestReverse", func() {
			head.Reverse()
			log.Println("after reverse :", head.Travel())
			So(head.Travel()[0], ShouldEqual, 3)
			So(head.Travel()[1], ShouldEqual, 2)
			So(head.Travel()[2], ShouldEqual, 1)
		})

	})
}
