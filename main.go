package main

import (
	"net/http"

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

	api := &binarytree.APIQueue{
		Store: &bt,
	}
	http.ListenAndServe(":8080", api)
}
