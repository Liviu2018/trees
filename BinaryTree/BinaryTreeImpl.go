package binarytree

import "math/rand"

type binaryTree struct {
	head *node
}

func (b *binaryTree) Create(value interface{}) BinaryTree {
	return b
}

func (b *binaryTree) Destroy() {
	b.head = nil
}

func (b *binaryTree) GetRoot() Node {
	return b.head
}

func (b *binaryTree) Add(value interface{}) {
	if b.head == nil {
		b.head = &node{value: value, left: nil, right: nil}

		return
	}

	b.addRecursive(b.head, value)
}

func (b *binaryTree) addRecursive(n *node, val interface{}) *node {
	if n == nil {
		n = &node{value: val, left: nil, right: nil}

		return n
	}

	left := rand.Intn(2)%2 == 0

	if left {
		n.left = b.addRecursive(n.left, val)
	} else {
		n.right = b.addRecursive(n.right, val)
	}

	return n
}

func (b *binaryTree) PreorderTraverse(Visit func(Node)) {
	b.preorderTraverseRec(b.head, Visit)
}

func (b *binaryTree) preorderTraverseRec(n *node, Visit func(Node)) {
	if n == nil {
		return
	}

	Visit(n)
	b.preorderTraverseRec(n.left, Visit)
	b.preorderTraverseRec(n.right, Visit)
}

func (b *binaryTree) InorderTraverse(Visit func(Node)) {
	b.inorderTraverseRec(b.head, Visit)
}

func (b *binaryTree) inorderTraverseRec(n *node, Visit func(Node)) {
	if n == nil {
		return
	}

	b.inorderTraverseRec(n.left, Visit)
	Visit(n)
	b.inorderTraverseRec(n.right, Visit)
}

func (b *binaryTree) PostorderTraverse(Visit func(Node)) {
	b.inorderTraverseRec(b.head, Visit)
}

func (b *binaryTree) postorderTraverseRec(n *node, Visit func(Node)) {
	if n == nil {
		return
	}

	b.postorderTraverseRec(n.left, Visit)
	b.postorderTraverseRec(n.right, Visit)
	Visit(n)
}

type node struct {
	value       interface{}
	left, right *node
}

func (n *node) LeftSubtree() Node {
	return n.left
}

func (n *node) RightSubtree() Node {
	return n.right
}

func (n *node) Value() interface{} {
	return n.value
}
