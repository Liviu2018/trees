package binarytree

import (
	"fmt"
	"math/rand"
	"strconv"
)

type BinaryTreeImpl struct {
	head *node
}

func (b *BinaryTreeImpl) Create(value interface{}) BinaryTree {
	return b
}

func (b *BinaryTreeImpl) Destroy() {
	b.head = nil
}

func (b *BinaryTreeImpl) GetRoot() Node {
	return b.head
}

func (b *BinaryTreeImpl) Add(value interface{}) {
	if b.head == nil {
		b.head = &node{value: value, left: nil, right: nil}

		return
	}

	b.addRecursive(b.head, value)
}

func (b *BinaryTreeImpl) addRecursive(n *node, val interface{}) *node {
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

func (b *BinaryTreeImpl) PreorderTraverse(Visit func(Node)) {
	b.preorderTraverseRec(b.head, Visit)
}

func (b *BinaryTreeImpl) preorderTraverseRec(n *node, Visit func(Node)) {
	if n == nil {
		return
	}

	Visit(n)
	b.preorderTraverseRec(n.left, Visit)
	b.preorderTraverseRec(n.right, Visit)
}

func (b *BinaryTreeImpl) InorderTraverse(Visit func(Node)) {
	b.inorderTraverseRec(b.head, Visit)
}

func (b *BinaryTreeImpl) inorderTraverseRec(n *node, Visit func(Node)) {
	if n == nil {
		return
	}

	b.inorderTraverseRec(n.left, Visit)
	Visit(n)
	b.inorderTraverseRec(n.right, Visit)
}

func (b *BinaryTreeImpl) PostorderTraverse(Visit func(Node)) {
	b.inorderTraverseRec(b.head, Visit)
}

func (b *BinaryTreeImpl) postorderTraverseRec(n *node, Visit func(Node)) {
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

// PrintByLevel prints the levels of a tree
func (b *BinaryTreeImpl) PrintByLevel() {
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
