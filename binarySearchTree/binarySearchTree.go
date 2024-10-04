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

func (curr *Node) deleteNodeByValueNew(value int) {
	if curr.isLeaf() {
		return
	}

	if value < curr.value {
		if curr.leftChild != nil {
			if value == curr.leftChild.value {
				// delete the left child node.
				tempRight := curr.leftChild.rightChild
				curr.leftChild = curr.leftChild.leftChild
				if tempRight != nil {
					curr.insertNode(tempRight)
				}
				return
			} else {
				curr.leftChild.deleteNodeByValueNew(value)
			}
		}
	} else {
		if curr.rightChild != nil {
			if value == curr.rightChild.value {
				// delete the right child
				tempRight := curr.rightChild.rightChild
				curr.rightChild = curr.rightChild.leftChild
				if tempRight != nil {
					curr.insertNode(tempRight)
				}
				return
			} else {
				curr.rightChild.deleteNodeByValueNew(value)
			}
		}
	}
}

func (bst *BinarySearchTree) deletenew(value int) {
	if bst.root == nil {
		return
	}

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

	// normal case
	bst.root.deleteNodeByValueNew(value)
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

func (curr *Node) deleteNodeFaster(value int) (delNode *Node) {
	if curr == nil {
		return nil
	}

	if value < curr.value {
		curr.leftChild = curr.leftChild.deleteNodeFaster(value)
	} else if value > curr.value {
		curr.rightChild = curr.rightChild.deleteNodeFaster(value)
	} else {
		if curr.leftChild == nil {
			return curr.rightChild
		} else if curr.rightChild == nil {
			return curr.leftChild
		}

		tempNode := curr.rightChild.getMinNode()
		curr.value = tempNode.value
		curr.rightChild.deleteNodeFaster(tempNode.value)
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

func (bst *BinarySearchTree) deleteFaster(value int) {
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

	_ = bst.root.deleteNodeFaster(value)
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
