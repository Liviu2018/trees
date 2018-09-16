package main

import (
	"fmt"

	binarytree "github.com/trees/BinaryTree/binary_tree"
)

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(vals)

	tree := binarytree.BinaryTreeImpl{}

	fmt.Println("\nadding values:")
	for _, v := range vals {
		tree.Add(v)
	}

	fmt.Println("\nInorderTraverse values:")
	tree.InorderTraverse(Visit)

	fmt.Println("\nPreorderTraverse values:")
	tree.PreorderTraverse(Visit)

	fmt.Println("\nPostorderTraverse values:")
	tree.PostorderTraverse(Visit)

	tree.PrintByLevel()
}

// Visit prints value of current node
func Visit(n binarytree.Node) {
	fmt.Print(n.Value())
}
