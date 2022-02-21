package binarytree

import (
	"strconv"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) Add(v int) {
	node := Node{Value: v}
	if bt.Root == nil {
		bt.Root = &node
		return
	}
	bt.Root.Insert(node)
}

func (n *Node) Insert(node Node) {
	if node.Value <= n.Value {
		if n.Left == nil {
			n.Left = &node
			return
		}
		n.Left.Insert(node)
		return
	}
	if n.Right == nil {
		n.Right = &node
		return
	}
	n.Right.Insert(node)
}

func (n Node) String() string {
	var msg string
	if n.Left != nil {
		msg = n.Left.String() + ","
	}

	msg += strconv.Itoa(n.Value)
	if n.Right != nil {
		msg += "," + n.Right.String()
	}
	return msg
}

func (bt BinaryTree) String() string {
	if bt.Root == nil {
		return "nil"
	}
	return bt.Root.String()
}

func (bt *BinaryTree) Path(v int) string {
	msg := "{"
	curr := bt.Root
	next := bt.Root
	for ; curr != nil && curr.Value != v; curr = next {

		if msg != "{" {
			msg += ", "
		}
		if v < curr.Value {
			msg += "L"
			next = curr.Left
			continue
		}
		if v > curr.Value {
			msg += "R"
			next = curr.Right
		}
	}
	if curr == nil {
		return "{Value not found}"
	}
	return msg + ", " + curr.String() + "}"
}

func (bt *BinaryTree) Max() string {
	msg := "{"
	for curr := bt.Root; curr != nil; curr = curr.Right {
		if curr.Right == nil {
			msg += curr.String()
		}
	}
	return msg + "}"
}

func (bt *BinaryTree) Min() string {
	msg := "{"
	for curr := bt.Root; curr != nil; curr = curr.Left {
		if curr.Left == nil {
			msg += curr.String()
		}
	}
	return msg + "}"
}

func (n *Node) Reverse() {
	if n == nil {
		return
	}
	n.Left.Reverse()
	n.Right.Reverse()
	n.Left, n.Right = n.Right, n.Left
}

func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	return n.Left.Size() + n.Right.Size() + 1
}

func (bt *BinaryTree) Size() string {
	s := bt.Root.Size()
	return "{Size = " + strconv.Itoa(s) + "}"
}

func (n *Node) MaxDepth() int {
	if n == nil {
		return 0
	}
	left := n.Left.MaxDepth() + 1
	right := n.Right.MaxDepth() + 1

	if left > right {
		return left
	}
	return right
}

func (bt *BinaryTree) MaxDepth() string {
	s := bt.Root.MaxDepth()
	return "{Max depth = " + strconv.Itoa(s) + "}"
}
