package main

import (
	"fmt"

	binarytree "github.com/deankinane/algo-course/src/data_structures/binary_tree"
)

func main() {
	root := binarytree.BinaryNode[int]{
		Val: 0,
	}

	FillTestTree(&root, 2)

	fmt.Print(root.PreOrderTraverse())
	fmt.Print(root.PostOrderTraverse())
}

var nodeIndex int = 0

func FillTestTree(node *binarytree.BinaryNode[int], depth int) {

	if depth < 0 {
		return
	}

	if node.Left == nil {
		nodeIndex++
		node.Left = &binarytree.BinaryNode[int]{
			Val: nodeIndex,
		}

		FillTestTree(node.Left, depth-1)
	}

	if node.Right == nil {
		nodeIndex++
		node.Right = &binarytree.BinaryNode[int]{
			Val: nodeIndex,
		}

		FillTestTree(node.Right, depth-1)
	}
}
