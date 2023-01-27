package stack

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLinkStack(t *testing.T) {
	Convey("TestLinkStack", t, func() {
		stack := new(LinkStack)
		stack.Push("car")
		stack.Push("city")
		stack.Push("cat")

		Convey("TestPush", func() {
			stack.Push("charge")
			So(stack.Size(), ShouldEqual, 4)
		})

		Convey("TestPop", func() {
			So(stack.Pop(), ShouldEqual, "cat")
		})

		Convey("TestPeek", func() {
			So(stack.Peek(), ShouldEqual, "cat")
		})

		Convey("TestSize", func() {
			So(stack.Size(), ShouldEqual, 3)
		})

		Convey("TestIsEmpty", func() {
			So(stack.IsEmpty(), ShouldBeFalse)
		})
	})
}
