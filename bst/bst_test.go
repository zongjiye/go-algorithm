package bst

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestBST(t *testing.T) {
	root := &BinarySearchTree{
		Val: 5,
		Left: &BinarySearchTree{
			Val: 3,
			Left: &BinarySearchTree{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &BinarySearchTree{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinarySearchTree{
			Val:  6,
			Left: nil,
			Right: &BinarySearchTree{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	Convey("TestBST", t, func() {
		Convey("Test IsBST", func() {
			So(root.IsBST(), ShouldBeTrue)
		})

		Convey("Test ToBST", func() {
			nums := []int{-10, -3, 0, 5, 9}
			bst := new(BinarySearchTree)
			bst.ToBST(nums)
			So(bst.Val, ShouldEqual, 0)
		})

		Convey("Test Search", func() {
			node := root.Search(6)
			So(node.Val, ShouldEqual, 6)
		})

		Convey("Test Search2", func() {
			node := root.Search2(9)
			So(node, ShouldBeNil)

			node1 := root.Search2(7)
			So(node1.Val, ShouldEqual, 7)
		})

		Convey("Test Delete", func() {
			log.Println(root.InOrderBST())
			root.Delete(3)
			log.Println(root.InOrderBST())
		})
	})

}
