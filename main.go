package main

import (
	"fmt"

	"github.com/matt-d-nelson/binarytree/binarytree"
)

func main() {
	var bt binarytree.BinaryTree
	bt.Add(10)
	bt.Add(5)
	bt.Add(4)
	bt.Add(7)
	bt.Add(17)
	bt.Add(15)
	bt.Add(3)

	fmt.Println(bt.Path(3))
	fmt.Println(bt.Path(18))

	bt.Root.Reverse()
	fmt.Println(bt.Path(3))

	bt.Root.Reverse()
	fmt.Println(bt.Path(3))
	fmt.Println(bt.Size())
	fmt.Println(bt.Max())
	fmt.Println(bt.MaxDepth())
}
