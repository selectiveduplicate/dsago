package binaryTree

import (
	"fmt"
	"io"
)

// TreeNode is a node in a tree
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

// CreateNode creates a node
func CreateNode(item int) *TreeNode {
	newNode := TreeNode{}
	newNode.data = item
	newNode.left = nil
	newNode.right = nil
	return &newNode
}

// AddLeft adds a left child
func (n *TreeNode) AddLeft(l *TreeNode) {
	n.left = l
}

// AddRight adds a right child
func (n *TreeNode) AddRight(r *TreeNode) {
	n.right = r
}

// PreOrder does the preorder traversal
func PreOrder(root *TreeNode, writer io.Writer) {
	fmt.Fprintf(writer, "%d", root.data)
	if root.left != nil {
		PreOrder(root.left, writer)
	}
	if root.right != nil {
		PreOrder(root.right, writer)
	}
}

// PostOrder does the postorder traversal
func PostOrder(root *TreeNode, writer io.Writer) {
	if root.left != nil {
		PostOrder(root.left, writer)
	}
	if root.right != nil {
		PostOrder(root.right, writer)
	}
	fmt.Fprintf(writer, "%d", root.data)
}

//InOrder does the inorder traversal
func InOrder(root *TreeNode, writer io.Writer) {
	if root.left != nil {
		InOrder(root.left, writer)
	}
	fmt.Fprintf(writer, "%d", root.data)
	if root.right != nil {
		InOrder(root.right, writer)
	}
}
