package binarysearchtree

import (
	"fmt"
	"strconv"
)

// BinarySearchTreeImpl is an implementation of binary tree
type BinarySearchTreeImpl struct {
	head    *node
	compare func(interface{}, interface{}) int
}

// Create creates and returns an empty tree
func (b BinarySearchTreeImpl) Create(compare func(interface{}, interface{}) int) BinarySearchTree {
	result := BinarySearchTreeImpl{head: nil, compare: compare}

	return result
}

// AddElement add a new element to our tree
func (b BinarySearchTreeImpl) AddElement(val interface{}) {
	if b.head == nil {
		b.head = &node{value: val, left: nil, right: nil, count: 1}

		return
	}

	b.addElementRecursive(b.head, val)
}

func (b BinarySearchTreeImpl) addElementRecursive(n *node, val interface{}) *node {
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

// RemoveElement removes an element from the binary search tree
func (b BinarySearchTreeImpl) RemoveElement(val interface{}) interface{} {
	if b.head == nil {
		return nil
	}

	if b.compare(b.head.value, val) == 0 {
		b.head.count--

		if b.head.count == 0 {
			b.head = b.deleteHeadAndBalanceSubtrees(b.head)
		}

		return val
	}

	return b.removeElementRecursive(b.head, val)
}

func (b BinarySearchTreeImpl) removeElementRecursive(n *node, val interface{}) interface{} {
	if b.compare(n.value, val) < 0 {
		if n.left == nil {
			return nil
		}

		if b.compare(n.left.value, val) == 0 {
			n.left.count--

			if n.left.count == 0 {
				n.left = b.deleteHeadAndBalanceSubtrees(n.left)
			}

			return val
		}

		return b.removeElementRecursive(n.left, val)
	}

	if n.right == nil {
		return nil
	}

	if b.compare(n.right.value, val) == 0 {
		n.right.count--

		if n.right.count == 0 {
			n.right = b.deleteHeadAndBalanceSubtrees(n.right)
		}

		return val
	}

	return b.removeElementRecursive(n.right, val)
}

func (b BinarySearchTreeImpl) deleteHeadAndBalanceSubtrees(n *node) *node {
	if n == nil {
		return nil
	}

	if n.right == nil {
		return n.left
	}

	if n.left == nil {
		return n.right
	}

	min := b.findMin(n.right)
	n.value = min.value

	if n.right.left == nil {
		n.right.count--

		if n.right.count == 0 {
			n.right = n.right.right
		}
	} else {
		b.removeMinRecursive(n.right)
	}

	return n
}

// RemoveAllOccurrences removes all occurences of a given value
func (b BinarySearchTreeImpl) RemoveAllOccurrences(val interface{}) int {
	if b.head == nil {
		return 0
	}

	if b.compare(b.head.value, val) == 0 {
		count := b.head.count
		b.head = b.deleteHeadAndBalanceSubtrees(b.head)

		return count
	}

	return b.removeAllOccurrencesRecursive(b.head, val)
}

func (b BinarySearchTreeImpl) removeAllOccurrencesRecursive(n *node, val interface{}) int {
	if b.compare(val, n.value) < 0 {
		if n.left == nil {
			return 0
		}

		if b.compare(n.left.value, val) == 0 {
			count := n.left.count

			n.left = b.deleteHeadAndBalanceSubtrees(n.left)

			return count
		}

		return b.removeAllOccurrencesRecursive(n.left, val)
	}

	if n.right == nil {
		return 0
	}

	if b.compare(n.right.value, val) == 0 {
		count := n.right.count

		n.right = b.deleteHeadAndBalanceSubtrees(n.right)

		return count
	}

	return b.removeAllOccurrencesRecursive(n.right, val)
}

// RemoveMin removes the minimum node out of this BST
func (b BinarySearchTreeImpl) RemoveMin() interface{} {
	if b.head == nil {
		return nil
	}

	if b.head.left == nil {
		val := b.head.value

		b.head = b.head.left

		return val
	}

	return b.removeMinRecursive(b.head)
}

func (b BinarySearchTreeImpl) removeMinRecursive(n *node) *node {
	if n == nil {
		return nil
	}

	if n.left == nil {
		return nil
	}

	for n.left.left != nil {
		n = n.left
	}

	result := n.left
	n.left = nil

	return result
}

// RemoveMax removes max value in BST
func (b BinarySearchTreeImpl) RemoveMax() interface{} {
	if b.head == nil {
		return nil
	}

	if b.head.left == nil {
		val := b.head.value

		b.head = b.head.left

		return val
	}

	return nil
}

func (b BinarySearchTreeImpl) removeMaxRecursive(n *node) *node {
	if n == nil {
		return nil
	}

	if n.right == nil {
		return nil
	}

	for n.right.right != nil {
		n = n.right
	}

	result := n.right
	n.right = nil

	return result
}

// FindMin return the min in a BST
func (b BinarySearchTreeImpl) FindMin() interface{} {
	return b.findMin(b.head)
}

func (b BinarySearchTreeImpl) findMin(n *node) *node {
	if n == nil {
		return nil
	}

	for n.left != nil {
		n = n.left
	}

	return n
}

// FindMax finds the max in this BST
func (b BinarySearchTreeImpl) FindMax() interface{} {
	return b.findMax(b.head)
}

func (b BinarySearchTreeImpl) findMax(n *node) *node {
	if n == nil {
		return nil
	}

	for n.right != nil {
		n = n.right
	}

	return n
}

// PreorderTraverse traverses tree in preorder
func (b BinarySearchTreeImpl) PreorderTraverse(Visit func(Node)) {
	b.preorderTraverseRec(b.head, Visit)
}

func (b BinarySearchTreeImpl) preorderTraverseRec(n *node, Visit func(Node)) {
	if n == nil {
		return
	}

	Visit(n)
	b.preorderTraverseRec(n.left, Visit)
	b.preorderTraverseRec(n.right, Visit)
}

// InorderTraverse traverses tree in order
func (b BinarySearchTreeImpl) InorderTraverse(Visit func(Node)) {
	b.inorderTraverseRec(b.head, Visit)
}

func (b BinarySearchTreeImpl) inorderTraverseRec(n *node, Visit func(Node)) {
	if n == nil {
		return
	}

	b.inorderTraverseRec(n.left, Visit)
	Visit(n)
	b.inorderTraverseRec(n.right, Visit)
}

// PostorderTraverse traverses tree in postorder
func (b BinarySearchTreeImpl) PostorderTraverse(Visit func(Node)) {
	b.inorderTraverseRec(b.head, Visit)
}

func (b BinarySearchTreeImpl) postorderTraverseRec(n *node, Visit func(Node)) {
	if n == nil {
		return
	}

	b.postorderTraverseRec(n.left, Visit)
	b.postorderTraverseRec(n.right, Visit)
	Visit(n)
}

// PrintByLevel prints the levels of a tree
func (b BinarySearchTreeImpl) PrintByLevel() {
	fmt.Println("printing tree")

	queue := []node{}
	queue = append(queue, *b.head)
	level := 0

	for len(queue) > 0 {
		fmt.Println()
		fmt.Println("level " + strconv.Itoa(level) + " :")

		temp := []node{}

		for _, v := range queue {
			if v.left != nil {
				temp = append(temp, *v.left)
			}

			if v.right != nil {
				temp = append(temp, *v.right)
			}

			fmt.Print(v.value)
			fmt.Print(" ")
		}

		queue = temp
		level++
	}
}

type node struct {
	value       interface{}
	count       int
	left, right *node
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

func (n *node) Value() interface{} {
	return n.value
}
