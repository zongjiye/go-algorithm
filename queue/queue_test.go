package queue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQueue(t *testing.T) {
	q := new(Queue)
	q.EnQueue("a")
	q.EnQueue("b")
	q.EnQueue("c")

	Convey("TestQueue", t, func() {
		Convey("TestEnqueue", func() {
			q.EnQueue("d")
			So(q.Size(), ShouldEqual, 4)
		})

		Convey("TestDequeue", func() {
			So(q.DeQueue(), ShouldEqual, "a")
		})

		Convey("TestFront", func() {
			So(q.Front(), ShouldEqual, "b")
		})

		Convey("TestIsEmpty", func() {
			So(q.IsEmpty(), ShouldBeFalse)
		})

		Convey("TestSize", func() {
			So(q.Size(), ShouldEqual, 3)
		})
	})
}
