package tree

import (
	"container/list"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestTree(t *testing.T) {
	list.New()
	Convey("TestTree", t, func() {
		root := NewTreeNode("A")
		// A<-B A->C
		root.SetLeft(NewTreeNode("B"))
		root.SetRight(NewTreeNode("C"))
		// B<-D B->E
		root.Left().SetLeft(NewTreeNode("D"))
		root.Left().SetRight(NewTreeNode("E"))
		// C<-F
		root.Right().SetLeft(NewTreeNode("F"))

		Convey("TestPreOrder", func() {
			So(strings.Join(root.PreOrder(), ""), ShouldEqual, "ABDECF")
		})

		Convey("TestInOrder", func() {
			So(strings.Join(root.InOrder(), ""), ShouldEqual, "DBEAFC")
		})

		Convey("TestPostOrder", func() {
			So(strings.Join(root.PostOrder(), ""), ShouldEqual, "DEBFCA")
		})

		Convey("TestPreOrderWithRecursion", func() {
			So(strings.Join(root.PreOrderWithRecursion(), ""), ShouldEqual, "ABDECF")
		})

		Convey("TestInOrderWithRecursion", func() {
			So(strings.Join(root.InOrderWithRecursion(), ""), ShouldEqual, "DBEAFC")
		})

		Convey("TestPostOrderWithRecursion", func() {
			So(strings.Join(root.PostOrderWithRecursion(), ""), ShouldEqual, "DEBFCA")
		})

		Convey("TestLayerOrder", func() {
			So(strings.Join(root.LayerOrder(), ""), ShouldEqual, "ABCDEF")
		})

		Convey("TestDFSOrder", func() {
			So(strings.Join(root.DFSOrder(), ""), ShouldEqual, "ABDECF")
		})
	})
}
