package main

import "fmt"
import "binaryTree"

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8}

	tree := binaryTree.binaryTree{head: nil}

	for _, v := range vals {
		tree.Add(v)
	}

	tree.PreorderTraverse()
}

func Visit(n binaryTree.Node) {
	fmt.Print(n.Value())
}
