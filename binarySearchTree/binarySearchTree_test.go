package main

import (
	"math/rand/v2"
	"testing"
)

func TestBstCreation(t *testing.T) {
	bst := &BinarySearchTree{root: nil}
	if bst.root != nil {
		t.Fatal("Bst root is not nil")
	}
}

func Test100BstInserts(t *testing.T) {
	bst := &BinarySearchTree{root: nil}

	for i := 0; i < 100; i++ {
		bst.insert(rand.IntN(100))
	}
}

func Test100RandomActions(t *testing.T) {
	bst := &BinarySearchTree{root: nil}

	for i := 0; i < 300000; i++ {
		rand_choice := rand.IntN(2)
		if rand_choice == 0 {
			bst.insert(rand.IntN(1000))
		} else {
			bst.deletenew(rand.IntN(1000))
		}
	}
	bst.inOrderTraverse()
}
