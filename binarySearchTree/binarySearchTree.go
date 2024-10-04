package main

import "fmt"

// Node: Each node in the tree
type Node struct {
	value      int
	leftChild  *Node
	rightChild *Node
}

// BinarySearchTree: We need only the root.
type BinarySearchTree struct {
	root *Node
}

// insertNode 'Insert a node starting at the root node.'
func (curr *Node) insertNode(newNode *Node) {
	if newNode == nil {
		return
	}

	if newNode.value < curr.value {
		if curr.leftChild == nil {
			curr.leftChild = newNode
		} else {
			curr.leftChild.insertNode(newNode)
		}
	} else {
		if curr.rightChild == nil {
			curr.rightChild = newNode
		} else {
			curr.rightChild.insertNode(newNode)
		}
	}
}

// insert - Insert a value into the bst.
func (bst *BinarySearchTree) insert(value int) {
	// create the new node
	newNode := Node{value: value, leftChild: nil, rightChild: nil}

	// if the bst is empty. Then just add this node to the root and done.
	if bst.root == nil {
		bst.root = &newNode
		return
	}

	// Bst already has some nodes.
	// insert the node at the correct position
	bst.root.insertNode(&newNode)
}

func (curr *Node) isLeaf() bool {
	if curr.leftChild == nil && curr.rightChild == nil {
		return true
	} else {
		return false
	}
}

func (curr *Node) deleteNode(value int) (delNode *Node) {
	if curr == nil {
		return nil
	}

	if value < curr.value {
		curr.leftChild = curr.leftChild.deleteNode(value)
	} else if value > curr.value {
		curr.rightChild = curr.rightChild.deleteNode(value)
	} else {
		if curr.leftChild == nil {
			return curr.rightChild
		} else if curr.rightChild == nil {
			return curr.leftChild
		}

		tempNode := curr.rightChild.getMinNode()
		curr.value = tempNode.value
		curr.rightChild.deleteNode(tempNode.value)
	}

	return curr
}

func (curr *Node) getMinNode() *Node {
	minNode := curr
	for minNode.leftChild != nil {
		minNode = minNode.leftChild
	}
	return minNode
}

func (bst *BinarySearchTree) delete(value int) {
	if bst.root == nil {
		fmt.Println("Cannot delete from an empty BST!!")
		return
	}

	// if the root is the only node and it is the one to be deleted.
	if bst.root.value == value {
		if bst.root.isLeaf() {
			bst.root = nil
			return
		}
	}

	_ = bst.root.deleteNode(value)
}

// inOrder: Node reciver for inOrder traversal.
func (curr *Node) inOrder() {
	if curr == nil {
		return
	}

	// left - parent - right ... right??
	curr.leftChild.inOrder()
	fmt.Println(curr.value)
	curr.rightChild.inOrder()
}

// inOrderTraverse: Traverse the BST in order.
func (bst BinarySearchTree) inOrderTraverse() {
	if bst.root == nil {
		fmt.Println("BST is empty!!")
		return
	}

	fmt.Println("InOrderTraverse---")
	// traverse in order
	bst.root.inOrder()
}

func main() {
	fmt.Println("A Binary Search Tree implementation.")

	// Create a new BST
	bst := &BinarySearchTree{root: nil}

	// try printing an empty bst,
	bst.inOrderTraverse()

	// insert a few nodes into the bst
	bst.insert(20)
	bst.insert(10)
	bst.insert(30)
	bst.insert(5)
	bst.insert(25)
	bst.insert(31)
	bst.delete(10)
	bst.inOrderTraverse()
	bst.delete(30)
	bst.inOrderTraverse()

	// Print the BST.
	bst.inOrderTraverse()
}
