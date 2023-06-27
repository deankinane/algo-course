package binarytree

import (
	"fmt"
	"testing"
)

func FillTestTree(node *BinaryNode[int], depth int, index *int) {

	if depth < 0 {
		return
	}

	if node.Left == nil {
		*index++
		node.Left = &BinaryNode[int]{
			Val: *index,
		}

		FillTestTree(node.Left, depth-1, index)
	}

	if node.Right == nil {
		*index++
		node.Right = &BinaryNode[int]{
			Val: *index,
		}

		FillTestTree(node.Right, depth-1, index)
	}
}

func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	index := 0
	result := "&[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]"

	root := BinaryNode[int]{
		Val: index,
	}

	FillTestTree(&root, 2, &index)

	data := root.PreOrderTraverse()

	if len(*data) != 15 {
		t.Errorf("Expected 15, got %v", len(*data))
	}

	if fmt.Sprint(data) != result {
		t.Errorf("Exptected %v, got %v", result, data)
	}
}
