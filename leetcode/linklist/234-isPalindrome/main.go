package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("val: %d, next node: %v", l.Val, l.Next)
}

// 时间复杂度O(n), 空间复杂度O(n), 最简单最弱智的方法
func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}

	return true
}

// reverse
func reverse(head *ListNode) *ListNode {
	pre := new(ListNode)
	p := head
	for p != nil {
		s := p
		p = p.Next
		s.Next = pre.Next
		pre.Next = s
	}

	return pre.Next
}

// endOfFirstHalf
func endOfFirstHalf(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 找到中间节点(1,2)
	firstHalf := endOfFirstHalf(head)
	// 逆转后半部分节点(2,1) -> (1,2)
	secondHalf := reverse(firstHalf.Next)
	// 进行比对
	for p, q := head, secondHalf; p != nil && q != nil; p, q = p.Next, q.Next {
		if p.Val != q.Val {
			return false
		}
	}

	// 复原这一波 如果需要用到原来的数据可以复原，如果不需要可以不复原
	//firstHalf.Next = reverse(secondHalf)
	return true
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	fmt.Println(isPalindrome2(head))
	fmt.Println(head.String())

	fmt.Println(isPalindrome(head))
	fmt.Println(head.String())

}
