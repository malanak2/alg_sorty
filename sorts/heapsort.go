package sorts

import (
	"container/list"
	"fmt"
	"io/fs"

	"github.com/malanak2/alg_sorts/util"
)

type Node struct {
	value any

	left  *Node
	right *Node
}

// // gemini print func
func PrintTree(root *Node) {
	// The root doesn't have a side, so we pass "ROOT"
	printRecursive(root, "", true, "ROOT")
}

func printRecursive(node *Node, prefix string, isLast bool, side string) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	// Format the label to show the side (L, R, or ROOT)
	label := fmt.Sprintf("[%s] ", side)

	if node == nil {
		fmt.Printf("%s%s%sNODE:NIL\n", prefix, connector, label)
		return
	}

	// Print the current node with its side label
	fmt.Printf("%s%s%s%v\n", prefix, connector, label, node.value)

	// Update prefix for children
	newPrefix := prefix
	if isLast {
		newPrefix += "    "
	} else {
		newPrefix += "│   "
	}

	// Recursively print Right then Left
	// We pass "R" and "L" to explicitly track the side
	printRecursive(node.right, newPrefix, false, "R")
	printRecursive(node.left, newPrefix, true, "L")
}

// End gemini

func HeapSort(file fs.DirEntry) {
	data, err := util.LoadFile(file)
	if err != nil {
		panic(err)
	}

	root := heapify(data.Front(), 0)
	PrintTree(root)
}

func sort(el *Node) {
	if el == nil {
		return
	}
	sort(el.left)
	sort(el.right)
	if el.left == nil || el.right == nil {
		return
	}

}

func heapify(data *list.Element, depth int) *Node {
	if data == nil {
		return nil
	}
	if data.Next() != nil {
		// Left is 2i + 1, right is 2i + 2
		index := depth
		el := data
		for index != 2*depth+1 {
			if el.Next() == nil {
				return &Node{value: data.Value, left: nil, right: nil}
			}
			el = el.Next()
			index++
		}
		return &Node{value: data.Value, left: heapify(el, 2*depth+1), right: heapify(el.Next(), 2*depth+2)}
	}
	return &Node{value: data.Value, left: nil, right: nil}
}
