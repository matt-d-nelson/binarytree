package binarytree

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
