package main

import "fmt"

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
func PreOrder(root *TreeNode) {
	fmt.Println(root.data)
	if root.left != nil {
		PreOrder(root.left)
	}
	if root.right != nil {
		PreOrder(root.right)
	}
}

// PostOrder does the postorder traversal
func PostOrder(root *TreeNode) {
	if root.left != nil {
		PostOrder(root.left)
	}
	if root.right != nil {
		PostOrder(root.right)
	}
	fmt.Println(root.data)
}

//InOrder does the inorder traversal
func InOrder(root *TreeNode) {
	if root.left != nil {
		InOrder(root.left)
	}
	fmt.Println(root.data)
	if root.right != nil {
		InOrder(root.right)
	}
}

// CreateTree creates a tree
func CreateTree() *TreeNode {
	root := CreateNode(5)
	ten := CreateNode(10)
	twenty := CreateNode(20)
	root.AddLeft(ten)
	root.AddRight(twenty)

	seven := CreateNode(7)
	nine := CreateNode(9)
	six := CreateNode(6)
	eight := CreateNode(8)
	ten.AddLeft(seven)
	ten.AddRight(nine)
	twenty.AddLeft(six)
	twenty.AddRight(eight)

	return root
}

func main() {
	tree := CreateTree()
	PreOrder(tree)
	fmt.Println("")
	PostOrder(tree)
	fmt.Println("")
	InOrder(tree)
}
