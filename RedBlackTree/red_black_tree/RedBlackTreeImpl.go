package redblacktree

// CreateBST creates and returns an empty BST
func CreateBST(compare func(interface{}, interface{}) int) RedBlackTree {
	return (&Impl{nil, nil}).Create(compare)
}

// Impl is an implementation of binary tree
type Impl struct {
	head    *node
	compare func(interface{}, interface{}) int
}

// Create creates and returns an empty tree
func (b *Impl) Create(compare func(interface{}, interface{}) int) RedBlackTree {
	result := Impl{head: nil, compare: compare}

	return &result
}

// AddElement add a new element to our tree
func (b *Impl) AddElement(val interface{}) {
	if b.head == nil {
		b.head = &node{value: val, left: nil, right: nil, parent: nil, count: 1}

		return
	}

	n := b.addElementRecursive(b.head, val)
	colorTree(n)
}

func (b *Impl) addElementRecursive(n *node, val interface{}) *node {
	if n == nil {
		n = &node{value: val, left: nil, right: nil, count: 1}

		return n
	}

	if b.compare(val, n.value) == 0 {
		n.count++

		return n
	}

	if b.compare(val, n.value) < 0 {
		n.left = b.addElementRecursive(n.left, val)
	} else {
		n.right = b.addElementRecursive(n.right, val)
	}

	return n
}

func (b *Impl) colorTree(n *node) {
	n.red = true

	for n.parent != nil && n.parent.red {
		if n.parent.IsLeftChild() {
			b.rotateParentLeft(n)
		} else {
			b.rotateParentRight(n)
		}
	}
}

//	TODO: repeat the if with right and left exchanged
func (b *Impl) rotateParentRight(n *node) {
	uncle := n.parent.parent.left

	if uncle.red {
		n.parent.red = false
		uncle.red = false
		n.parent.parent.red = true

		n = n.parent.parent
	} else {
		if n.IsLeftChild() {
			n = n.parent

			b.rotateRight(n)
		}

		n.parent.red = false
		n.parent.parent.red = true

		b.rotateLeft(n.parent.parent)
	}
}

func (b *Impl) rotateParentLeft(n *node) {
	uncle := n.parent.parent.right

	if uncle.red {
		n.parent.red = false
		uncle.red = false
		n.parent.parent.red = true

		n = n.parent.parent
	} else {
		if !n.IsLeftChild() {
			n = n.parent

			b.rotateLeft(n)
		}

		n.parent.red = false
		n.parent.parent.red = true

		b.rotateRight(n.parent.parent)
	}
}

func (b *Impl) rotateLeft(n *node) {
	other := n.right

	other.parent = n.parent
	if n.parent == nil {
		b.head = other
	} else {
		if n.IsLeftChild() {
			n.parent.left = other
		} else {
			n.parent.right = other
		}
	}

	n.right = other.left
	if other.left != nil {
		other.left.parent = n
	}

	other.left = n
	n.parent = other
}

func (b *Impl) rotateRight(n *node) {
	if n.parent == nil {
		return
	}

	other := n.parent
	n.parent = other.parent

	if other.parent == nil {
		b.head = n
	} else {
		if other.IsLeftChild() {
			other.parent.left = n
		} else {
			other.parent.right = n
		}
	}

	other.left = n.right
	if other.left != nil {
		other.left.parent = other
	}

	n.right = other
	other.parent = n
}
