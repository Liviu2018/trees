package main

import "github.com/trees/BinarySearchTree/binary_search_tree"

func main() {
	tree := binarysearchtree.Create(compare)

	vals := []int{1, 22, 13, 4, 15, 6}
	for _, v := range vals {
		tree.AddElement(v, compare)
	}

	tree.PrintByLevel()
}

func compare(a, b interface{}) int {
	return a.(int) - b.(int)
}
