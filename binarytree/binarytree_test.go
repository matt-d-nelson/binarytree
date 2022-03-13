package binarytree

import "testing"

func assertIndexValue(t *testing.T, bt BinaryTree, idx int, expected int) {
	val, err := bt.Index(idx)
	if err != nil {
		t.Fatalf("Index of %d errored: %v", idx, err)
	}
	if val != expected {
		t.Fatalf("Index of %d = %d; want %v", idx, val, expected)
	}
}

func TestIndex0(t *testing.T) {
	var bt BinaryTree
	assertIndexValue(t, bt, 0, -1)
	bt.Add(10)
	assertIndexValue(t, bt, 0, 10)
	bt.Add(5)
	assertIndexValue(t, bt, 0, 5)
	bt.Add(6)
	assertIndexValue(t, bt, 0, 5)
}

func TestIndexEvery(t *testing.T) {
	var bt BinaryTree
	bt.Add(10)
	bt.Add(7)
	bt.Add(5)
	bt.Add(6)
	bt.Add(8)
	bt.Add(20)
	bt.Add(15)
	bt.Add(16)
	assertIndexValue(t, bt, 0, 5)
	assertIndexValue(t, bt, 1, 6)
	assertIndexValue(t, bt, 2, 7)
	assertIndexValue(t, bt, 3, 8)
	assertIndexValue(t, bt, 4, 10)
	assertIndexValue(t, bt, 5, 15)
	assertIndexValue(t, bt, 6, 16)
	assertIndexValue(t, bt, 7, 20)
}
