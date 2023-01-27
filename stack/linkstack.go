package stack

import "sync"

type LinkStack struct {
	root *linkNode
	size int
	lock sync.Mutex
}

type linkNode struct {
	value string
	next  *linkNode
}

func (ls *LinkStack) Push(v string) {
	ls.lock.Lock()
	defer ls.lock.Unlock()

	if ls.root == nil {
		ls.root = new(linkNode)
		ls.root.value = v
	} else {
		node := new(linkNode)
		node.value = v
		// insert the new node into the head of the linked list
		preNode := ls.root
		node.next = preNode
		ls.root = node
	}
	ls.size++
}

func (ls *LinkStack) Pop() string {
	ls.lock.Lock()
	defer ls.lock.Unlock()

	if ls.size == 0 {
		panic("empty")
	}

	topNode := ls.root
	v := topNode.value

	ls.root = topNode.next

	ls.size--
	return v
}

func (ls *LinkStack) Peek() string {
	if ls.size == 0 {
		panic("empty")
	}

	return ls.root.value
}

func (ls *LinkStack) IsEmpty() bool {
	return ls.size == 0
}

func (ls *LinkStack) Size() int {
	return ls.size
}
