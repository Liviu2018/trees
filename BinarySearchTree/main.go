package main

import "github.com/trees/BinarySearchTree/binary_search_tree"

func main() {
	tree := binarysearchtree.CreateBST(compare)

	vals := []int{1, 22, 13, 4, 15, 6}
	for _, v := range vals {
		tree.AddElement(v)
	}

	tree.PrintByLevel()
}

func compare(a, b interface{}) int {
	return a.(int) - b.(int)
}
