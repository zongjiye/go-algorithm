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

func deleteNode(root *BinarySearchTree, val int) *BinarySearchTree {
	if root == nil {
		return nil
	}

	if val < root.Val {
		root.Left = deleteNode(root.Left, val)
		return root
	}

	if val > root.Val {
		root.Right = deleteNode(root.Right, val)
	}

	// 找到要删除的节点位置
	// 如果左子树为空 就返回右子树
	if root.Left == nil {
		return root.Right
	}

	if root.Right == nil {
		return root.Left
	}

	// 左右 孩子都不空 从右子树中找到最小的节点
	minNode := root.Right
	for minNode.Left != nil {
		minNode = minNode.Left
	}

	// 删除最小的节点
	root.Val = minNode.Val
	root.Right = deleteMinNode(root.Right)
	return root
}

func deleteMinNode(root *BinarySearchTree) *BinarySearchTree {
	// 找到最小的节点时, 删除右节点
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	// 一直向左找找到最小的左孩子
	root.Left = deleteMinNode(root.Left)
	return root
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
