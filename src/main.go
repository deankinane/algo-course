package main

import (
	"fmt"

	binarytree "github.com/deankinane/algo-course/src/data_structures/binary_tree"
)

func main() {
	n1 := binarytree.BinaryNode[int]{
		Val: 0,
	}

	n2 := binarytree.BinaryNode[int]{
		Val: 0,
	}

	FillTestTree(&n1, 2)
	nodeIndex = 0
	FillTestTree(&n2, 3)

	fmt.Print(binarytree.BinaryTreeComparison[int](&n1, &n2))
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
