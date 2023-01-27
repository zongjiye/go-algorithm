package tree

/*
Tree 树的定义：

	i.有节点间的层次关系，分为父节点和子节点。
	ii.有唯一一个根节点，该根节点没有父节点。
	iii.除了根节点，每个节点有且只有一个父节点。
	iv.每一个节点本身以及它的后代也是一棵树，是一个递归的结构。
	v.没有后代的节点称为叶子节点，没有节点的树称为空树。

BinaryTree

	i. 二叉树：每个节点最多只有两个儿子节点的树。

	ii. 满二叉树：叶子节点与叶子节点之间的高度差为 0 的二叉树，即整棵树是满的，
	树呈满三角形结构。在国外的定义，非叶子节点儿子都是满的,树就是满二叉树。我们以国内为准。

	iii. 完全二叉树：完全二叉树是由满二叉树而引出来的，设二叉树的深度为 k，除第 k 层外，
	其他各层的节点数都达到最大值，且第 k 层所有的节点都连续集中在最左边。

	其数学性质：
	1. 高度为 h≥0 的二叉树至少有 h+1 个结点，比如最不平衡的二叉树就是退化的线性链表结构，所有的节点都只有左儿子节点，或者所有的节点都只有右儿子节点。
	2. 高度为 h≥0 的二叉树至多有 2^h+1 个节点，比如这棵树是满二叉树。
	3. 含有 n≥1 个结点的二叉树的高度至多为 n-1，由 1 退化的线性链表可以反推。
	4. 含有 n≥1 个结点的二叉树的高度至少为 logn，由 2 满二叉树可以反推。  logn <= h <= n-1
	5. 在二叉树的第 i 层，至多有 2^(i-1) 个节点，比如该层是满的。
*/
type Tree struct {
	data  string
	left  *Tree
	right *Tree
}

func NewTreeNode(data string) *Tree {
	return &Tree{data: data}
}

func (t *Tree) Data() string {
	if t == nil {
		panic("pointer nil")
	}

	return t.data
}

func (t *Tree) SetData(data string) {
	if t == nil {
		panic("pointer nil")
	}
	t.data = data
}

func (t *Tree) Left() *Tree {
	return t.left
}

func (t *Tree) SetLeft(node *Tree) {
	node.left = t.left
	t.left = node

}

func (t *Tree) Right() *Tree {
	return t.right
}

func (t *Tree) SetRight(node *Tree) {
	node.right = t.right
	t.right = node
}

// PreOrder without recursion
func (t *Tree) PreOrder() []string {
	if t == nil {
		return []string{}
	}

	res := make([]string, 0)
	stack := make([]*Tree, 0)
	node := t
	for node != nil || len(stack) != 0 {
		if node != nil {
			res = append(res, node.data)
			stack = append(stack, node)
			node = node.left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = node.right
		}
	}

	return res
}

// InOrder without recursion
func (t *Tree) InOrder() []string {
	if t == nil {
		return []string{}
	}

	res := make([]string, 0)
	stack := make([]*Tree, 0)
	node := t

	for node != nil || len(stack) != 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.data)
			node = node.right
		}
	}

	return res
}

// PostOrder without recursion
func (t *Tree) PostOrder() []string {
	if t == nil {
		return []string{}
	}

	res := make([]string, 0)
	stack := []*Tree{t}
	var preNode *Tree
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		if (top.left == nil && top.right == nil) || (top != nil && (top.left == preNode || top.right == preNode)) {
			res = append(res, top.data)
			stack = stack[:len(stack)-1]
			// pre visited node
			preNode = top
		} else {
			if top.right != nil {
				stack = append(stack, top.right)
			}

			if top.left != nil {
				stack = append(stack, top.left)
			}
		}

	}
	return res
}

// LayerOrder BFS
func (t *Tree) LayerOrder() []string {
	if t == nil {
		return []string{}
	}

	res := make([]string, 0)
	queue := []*Tree{t}

	// while nodes length is 0
	for len(queue) != 0 {
		current := queue[0]
		res = append(res, current.data)
		queue = queue[1:]

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return res
}

// DFSOrder DFS like PreOrder
func (t *Tree) DFSOrder() []string {
	if t == nil {
		return []string{}
	}

	res := make([]string, 0)
	stack := []*Tree{t}

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		res = append(res, current.data)
		stack = stack[0 : len(stack)-1]
		if current.right != nil {
			stack = append(stack, current.right)
		}

		if current.left != nil {
			stack = append(stack, current.left)
		}
	}

	return res
}

func (t *Tree) PreOrderWithRecursion() []string {
	res := make([]string, 0)
	if t == nil {
		return res
	}

	var dfs func(tree *Tree)
	dfs = func(tree *Tree) {
		if tree == nil {
			return
		}
		res = append(res, tree.data)
		dfs(tree.left)
		dfs(tree.right)
	}

	dfs(t)
	return res
}

func (t *Tree) InOrderWithRecursion() []string {
	res := make([]string, 0)
	if t == nil {
		return res
	}

	var dfs func(tree *Tree)
	dfs = func(tree *Tree) {
		if tree == nil {
			return
		}
		dfs(tree.left)
		res = append(res, tree.data)
		dfs(tree.right)
	}

	dfs(t)
	return res
}

func (t *Tree) PostOrderWithRecursion() []string {
	res := make([]string, 0)
	if t == nil {
		return res
	}

	var dfs func(tree *Tree)
	dfs = func(tree *Tree) {
		if tree == nil {
			return
		}
		dfs(tree.left)
		dfs(tree.right)
		res = append(res, tree.data)
	}

	dfs(t)
	return res
}
