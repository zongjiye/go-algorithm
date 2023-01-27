package set

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet(10)
	s.Add(1)
	s.Add(199)

	Convey("TestSet", t, func() {
		Convey("TestAdd", func() {
			s.Add(10)
			So(s.Has(10), ShouldBeTrue)
		})

		Convey("TestRemove", func() {
			s.Remove(1)
			So(s.Has(1), ShouldBeFalse)
		})

		Convey("TestHas", func() {
			So(s.Has(199), ShouldBeTrue)
			So(s.Has(99), ShouldBeFalse)
		})

		Convey("TestLen", func() {
			So(s.Len(), ShouldEqual, 2)
		})

		Convey("TestIsEmpty And TestClear", func() {
			So(s.IsEmpty(), ShouldBeFalse)

		})

		Convey("TestToList", func() {
			log.Println(s.ToList())
		})

		Convey("TestClear", func() {
			s.Clear()
			So(s.IsEmpty(), ShouldBeTrue)
		})
	})
}
