package redblacktree

// RedBlackTree defines the ADT operations of a red black tree, almost the same as a nomary tree
type RedBlackTree interface {
	Create(compare func(interface{}, interface{}) int) RedBlackTree

	AddElement(interface{})
	RemoveElement(interface{}) interface{}

	RemoveAllOccurrences(interface{}) int
	RemoveMin() interface{}
	RemoveMax() interface{}

	FindMin() interface{}
	FindMax() interface{}

	PreorderTraverse(Visit func(Node))
	InorderTraverse(Visit func(Node))
	PostorderTraverse(Visit func(Node))

	PrintByLevel()
}

// Node is a node in the binary tree
type Node interface {
	LeftSubtree() Node
	RightSubtree() Node
	Parent() Node
	Color() bool
	IsLeftChild() bool
	Value() interface{}
}
