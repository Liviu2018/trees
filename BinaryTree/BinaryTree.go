package binarytree

// BinaryTree defines the ADT operations of a binary tree
type BinaryTree interface {
	Create(value interface{}) BinaryTree

	Destroy()
	GetRoot() Node

	Add(value interface{})

	PreorderTraverse(Visit func(Node))
	InorderTraverse(Visit func(Node))
	PostorderTraverse(Visit func(Node))
}

// Node is a node in the binary tree
type Node interface {
	LeftSubtree() Node
	RightSubtree() Node
	Value() interface{}
}
