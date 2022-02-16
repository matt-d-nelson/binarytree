package binarytree

import (
	"fmt"
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
		node.Left = nil
		node.Right = nil
		return
	} else {
		bt.Root.Insert(v)
	}
}

func (n *Node) Insert(v int) {
	node := Node{Value: v}
	if v <= n.Value {
		if n.Left == nil {
			n.Left = &node
			node.Left = nil
			node.Right = nil
		} else {
			n.Left.Insert(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &node
			node.Left = nil
			node.Right = nil
		} else {
			n.Right.Insert(v)
		}
	}
}

func (n Node) String() string {
	return fmt.Sprintf("%v", n.Value)
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
		} else if v > curr.Value {
			msg += "R"
			next = curr.Right
		}
	}
	if curr == nil {
		msg = "{Value not found"
	} else {
		msg += ", " + curr.String()
	}
	msg += "}"
	return msg
}

func (bt *BinaryTree) Max() string {
	msg := "{"
	for curr := bt.Root; curr != nil; curr = curr.Right {
		if curr.Right == nil {
			msg += curr.String()
		}
	}
	msg += "}"
	return msg
}

func (bt *BinaryTree) Min() string {
	msg := "{"
	for curr := bt.Root; curr != nil; curr = curr.Left {
		if curr.Left == nil {
			msg += curr.String()
		}
	}
	msg += "}"
	return msg
}

func (n *Node) Reverse() {
	if n == nil {
		return
	} else {
		n.Left.Reverse()
		n.Right.Reverse()

		temp := n.Left
		n.Left = n.Right
		n.Right = temp
	}
}

func (n *Node) Size() int {
	if n == nil {
		return 0
	} else {
		return (n.Left.Size() + n.Right.Size() + 1)
	}
}

func (bt *BinaryTree) Size() string {
	s := bt.Root.Size()
	return "{Size = " + strconv.Itoa(s) + "}"
}

func (n *Node) MaxDepth() int {
	if n == nil {
		return 0
	} else {
		left := n.Left.MaxDepth() + 1
		right := n.Right.MaxDepth() + 1

		if left > right {
			return left
		} else {
			return right
		}
	}
}

func (bt *BinaryTree) MaxDepth() string {
	s := bt.Root.MaxDepth()
	return "{Max depth = " + strconv.Itoa(s) + "}"
}
