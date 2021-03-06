package binarysearchtree

// BinarySearchTree has the ADT operations of a binary search tree
type BinarySearchTree interface {
	Create(compare func(interface{}, interface{}) int) BinarySearchTree

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
	Value() interface{}
}
