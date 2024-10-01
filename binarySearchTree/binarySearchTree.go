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

// deleteNodeByValue: Node receiver which recursively checks and deletes a node if matched.
func (curr *Node) deleteNodeByValue(value int, parent *Node) {
	if curr == nil {
		fmt.Println("Value not present.")
		return
	}
	// If curr node is to be deleted
	if curr.value == value {
		fmt.Println("Value found...deleting...")
		leftChild := curr.leftChild
		rightChild := curr.rightChild

		// We need to check if the current node is a left or right child.
		if curr.value < parent.value {
			if leftChild != nil {
				parent.leftChild = leftChild
				parent.leftChild.insertNode(rightChild)
			} else {
				parent.leftChild = rightChild
			}
			return
		} else {
			if leftChild != nil {
				parent.rightChild = leftChild
				parent.rightChild.insertNode(rightChild)
			} else {
				parent.rightChild = rightChild
			}
			return
		}
	} else if value < curr.value {
		curr.leftChild.deleteNodeByValue(value, curr)
	} else {
		curr.rightChild.deleteNodeByValue(value, curr)
	}
}

// delete: Delete a value from the BST if present.
func (bst *BinarySearchTree) delete(value int) {
	// If the BST is empty
	if bst.root == nil {
		fmt.Println("Cannot delete from an empty BST!!")
		return
	}

	fmt.Println("Searching for value -", value)

	// Edge case where the root is the node to be deleted.
	if bst.root.value == value {
		// if root is the only node
		if bst.root.isLeaf() {
			fmt.Println("Deleting root...")
			bst.root = nil
			return
		} else {
			leftChild := bst.root.leftChild
			rightChild := bst.root.rightChild

			if leftChild != nil {
				bst.root = leftChild
				bst.root.insertNode(rightChild)
			} else {
				bst.root = rightChild
			}
			return
		}
	}

	// normal case.
	if value < bst.root.value {
		bst.root.leftChild.deleteNodeByValue(value, bst.root)
	} else {
		bst.root.rightChild.deleteNodeByValue(value, bst.root)
	}
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
