package redblacktree

type node struct {
	value               interface{}
	count               int
	red                 bool
	left, right, parent *node
}

func (n *node) GetCount() int {
	return n.count
}

func (n *node) LeftSubtree() Node {
	return n.left
}

func (n *node) RightSubtree() Node {
	return n.right
}

func (n *node) Parent() Node {
	return n.parent
}

func (n *node) Color() bool {
	return n.red
}

func (n *node) IsLeftChild() bool {
	if n.parent == nil {
		return false
	}

	return n.parent.left == n
}

func (n *node) Value() interface{} {
	return n.value
}
