package stack

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStack(t *testing.T) {
	s := new(Stack)
	s.Push("a")
	s.Push("b")
	s.Push("c")
	Convey("TestStack", t, func() {
		Convey("TestPush", func() {
			s.Push("d")
			So(s.Size(), ShouldEqual, 4)
		})

		Convey("TestPop", func() {
			So(s.Pop(), ShouldEqual, "d")
		})

		Convey("TestPeek", func() {
			So(s.Peek(), ShouldEqual, "c")
		})

		Convey("TestIsEmpty", func() {
			So(s.IsEmpty(), ShouldBeFalse)
		})

		Convey("TestSize", func() {
			So(s.Size(), ShouldEqual, 3)
		})
	})
}
