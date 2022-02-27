package main

import (
	"net/http"

	"github.com/matt-d-nelson/binarytree/binarytree"
	"github.com/matt-d-nelson/linkedlist/linkedlist"
)

func main() {
	var bt binarytree.BinaryTree
	var ll linkedlist.LinkedList
	bt.Add(10)
	bt.Add(5)
	bt.Add(4)
	bt.Add(7)
	bt.Add(17)
	bt.Add(15)
	bt.Add(3)
	ll.Add(12)

	api := &binarytree.APIQueue{
		Store: &ll,
	}
	http.ListenAndServe(":8080", api)
}
