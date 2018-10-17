package binarytree_test

import (
	"fmt"
	"strconv"
	"testing"

	binarytree "github.com/trees/BinaryTree/binary_tree"
)

func BenchmarkAddNodes(b *testing.B) {
	tree := binarytree.BinaryTreeImpl{}

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		tree.Add(n)
	}

	size := 0
	f := func(n binarytree.Node) {
		size++
	}

	tree.InorderTraverse(f)

	fmt.Println("result size = " + strconv.Itoa(size))
}
