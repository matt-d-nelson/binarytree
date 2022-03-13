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

// func TestIndexEvery(t *testing.T) {
// 	var bt BinaryTree
// 	bt.Add(10)
// 	bt.Add(7)
// 	bt.Add(5)
// 	bt.Add(6)
// 	bt.Add(8)
// 	bt.Add(20)
// 	bt.Add(15)
// 	bt.Add(16)
// 	if bt.Index(0) != 5 {
// 		t.Fatalf("Index of 0 = %d; want 5", bt.Index(0))
// 	}
// 	if bt.Index(1) != 6 {
// 		t.Fatalf("Index of 1 = %d; want 6", bt.Index(1))
// 	}
// 	if bt.Index(2) != 7 {
// 		t.Fatalf("Index of 2 = %d; want 7", bt.Index(2))
// 	}
// 	if bt.Index(3) != 8 {
// 		t.Fatalf("Index of 3 = %d; want 8", bt.Index(3))
// 	}
// 	if bt.Index(4) != 10 {
// 		t.Fatalf("Index of 0 = %d; want 10", bt.Index(4))
// 	}
// 	if bt.Index(5) != 15 {
// 		t.Fatalf("Index of 5 = %d; want 15", bt.Index(5))
// 	}
// 	if bt.Index(6) != 16 {
// 		t.Fatalf("Index of 0 = %d; want 16", bt.Index(6))
// 	}
// 	if bt.Index(7) != 20 {
// 		t.Fatalf("Index of 7 = %d; want 20", bt.Index(7))
// 	}
// }
