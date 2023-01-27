package queue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDeque(t *testing.T) {
	q := new(Deque)
	q.AddFront("curry")
	q.AddFront("dream")
	q.AddRear("klay")
	q.AddRear("durant")

	Convey("TestQueue", t, func() {
		Convey("TestAddFront", func() {
			q.AddFront("lucky")
			So(q.RemoveFront(), ShouldEqual, "lucky")
		})

		Convey("TestAddRear", func() {
			q.AddRear("sasuke")
			So(q.RemoveRear(), ShouldEqual, "sasuke")
		})

		Convey("TestRemoveFront", func() {
			So(q.RemoveFront(), ShouldEqual, "dream")
		})

		Convey("TestRemoveRear", func() {
			So(q.RemoveRear(), ShouldEqual, "durant")
		})

		Convey("TestSize", func() {
			So(q.Size(), ShouldEqual, 2)
		})

		Convey("TestIsEmpty", func() {
			So(q.IsEmpty(), ShouldBeFalse)
		})
	})
}
