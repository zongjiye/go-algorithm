package linklist

type LinkNode struct {
	Data     int
	NextNode *LinkNode
}

func NewLinkList() *LinkNode {
	// The head node uses the data to
	// record the length of the link list
	return &LinkNode{
		Data:     0,
		NextNode: nil,
	}
}

func (l *LinkNode) Insert(pos int, node *LinkNode) {
	if pos-1 < 0 || pos-1 > l.Length() {
		panic("insertion position error")
	}

	p := l
	for i := 0; i < pos-1; i++ {
		p = p.NextNode
	}

	node.NextNode = p.NextNode
	p.NextNode = node
	l.Data++
}

func (l *LinkNode) Delete(pos int) {
	if pos-1 < 0 || pos-1 > l.Length() {
		panic("delete position error")
	}

	p := l
	for i := 0; i < pos-1; i++ {
		p = p.NextNode
	}

	s := p.NextNode
	p.NextNode = s.NextNode
	l.Data--
}

func (l *LinkNode) Length() int {
	return l.Data
}

func (l *LinkNode) Travel() []int {
	nodes := make([]int, 0)
	p := l.NextNode
	for p != nil {
		nodes = append(nodes, p.Data)
		p = p.NextNode
	}

	return nodes
}

func (l *LinkNode) HeadInsertion(array []int) {
	for _, e := range array {
		p := new(LinkNode)
		p.Data = e

		p.NextNode = l.NextNode
		l.NextNode = p
		l.Data++
	}
}

func (l *LinkNode) TailInsertion(array []int) {
	p := l
	for _, e := range array {
		node := new(LinkNode)
		node.Data = e

		p.NextNode = node
		p = p.NextNode
		l.Data++
	}
}

func (l *LinkNode) Reverse() {
	p := l.NextNode
	l.NextNode = nil
	for p != nil {
		// current node
		s := p
		// move pointer
		p = p.NextNode
		// head insertion
		s.NextNode = l.NextNode
		l.NextNode = s

	}
}
