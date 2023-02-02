package bst

import "math"

/*
	二叉搜索数的三个特性：
	1.中序遍历是递增序列
	2.二叉搜索树中的每一个节点的左子树都比当前节点值小
	3.二叉搜索树中的每一个节点的右子树都比当前节点值大
*/

// BinarySearchTree 二叉搜索树
type BinarySearchTree struct {
	Val   int
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

func (bst *BinarySearchTree) IsBST() bool {
	var isBST func(root *BinarySearchTree, min, max int) bool
	isBST = func(root *BinarySearchTree, min, max int) bool {
		if root == nil {
			return true
		}

		if min >= root.Val || max <= root.Val {
			return false
		}

		return isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max)
	}

	return isBST(bst, math.MinInt64, math.MaxInt64)
}

func (bst *BinarySearchTree) ToBST(nums []int) {
	var helper func(nums []int, left, right int) *BinarySearchTree
	helper = func(nums []int, left, right int) *BinarySearchTree {
		if left > right {
			return nil
		}
		// 总是选择中间位置左边的数字作为根节点
		mid := (left + right) / 2
		root := &BinarySearchTree{Val: nums[mid]}
		root.Left = helper(nums, left, mid-1)
		root.Right = helper(nums, mid+1, right)
		return root
	}

	*bst = *helper(nums, 0, len(nums)-1)
}

func (bst *BinarySearchTree) Search(val int) *BinarySearchTree {
	if bst == nil {
		return nil
	}

	if bst.Val == val {
		return bst
	} else if val < bst.Val {
		// val小于当前节点值, 在左子树找
		return bst.Left.Search(val)
	} else {
		// val大于当前节点值, 在右子树找
		return bst.Right.Search(val)
	}

}

func (bst *BinarySearchTree) Search2(val int) *BinarySearchTree {
	if bst == nil {
		return nil
	}

	root := bst
	for root != nil {
		if root.Val == val {
			return root
		} else if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return nil
}

func (bst *BinarySearchTree) Delete(val int) {
	*bst = *deleteNode(bst, val)
}

func (bst *BinarySearchTree) InOrderBST() (res []int) {
	res = make([]int, 0)
	var order func(root *BinarySearchTree)
	order = func(root *BinarySearchTree) {
		if root == nil {
			return
		}
		order(root.Left)
		res = append(res, root.Val)
		order(root.Right)
	}
	order(bst)
	return res
}

func deleteNode(root *BinarySearchTree, key int) *BinarySearchTree {
	switch {
	case root == nil:
		return nil
	case key < root.Val:
		root.Left = deleteNode(root.Left, key)
	case key > root.Val:
		root.Right = deleteNode(root.Right, key)
	case root.Left == nil || root.Right == nil:
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	default:
		// 从右子树中寻找最小的节点
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		//
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}
