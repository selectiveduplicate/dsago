package binaryTree

import (
	"bytes"
	"reflect"
	"testing"
)

// Create a tree for test
func CreateTestTree() *TreeNode {
	//		   1
	//       /   \
	//      2     3
	//    / \    / \
	//   7   9  6   8
	//	/ \			\
	//  5 4			10
	//
	//

	root := CreateNode(1)
	two := CreateNode(2)
	three := CreateNode(3)
	root.AddLeft(two)
	root.AddRight(three)

	seven := CreateNode(7)
	nine := CreateNode(9)
	six := CreateNode(6)
	eight := CreateNode(8)
	five := CreateNode(5)
	four := CreateNode(4)
	ten := CreateNode(10)

	two.AddLeft(seven)
	two.AddRight(nine)
	three.AddLeft(six)
	three.AddRight(eight)

	seven.AddLeft(five)
	seven.AddRight(four)
	eight.AddRight(ten)

	return root
}

func TestCreateNode(t *testing.T) {
	got := CreateNode(100)
	want := &TreeNode{100, nil, nil}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("creating node failed, wanted %v got %v", want, got)
	}
}

func TestAddLeftAndRight(t *testing.T) {
	parent := CreateNode(100)
	left := CreateNode(25)
	right := CreateNode(66)

	t.Run("adds a left child to parent node", func(t *testing.T) {

		parent.AddLeft(left)

		if parent.left != left {
			t.Errorf("failed to add left child to node %v; wanted %v got %v",
				parent, left, parent.left)
		}
	})
	t.Run("adds a right child to parent node", func(t *testing.T) {
		parent.AddRight(right)

		if parent.right != right {
			t.Errorf("failed to add right child to node %v; wanted %v got %v",
				parent, right, parent.right)
		}
	})
}

func TestPreOrder(t *testing.T) {
	tree := CreateTestTree()

	buffer := bytes.Buffer{}
	PreOrder(tree, &buffer)

	got := buffer.String()
	want := "12754936810"

	if got != want {
		t.Errorf("pre-order traversal failed, wanted %v got %v", want, got)
	}
}

func TestPostOrder(t *testing.T) {
	tree := CreateTestTree()

	buffer := bytes.Buffer{}
	PostOrder(tree, &buffer)

	got := buffer.String()
	want := "54792610831"

	if got != want {
		t.Errorf("post-order traversal failed, wanted %v got %v", want, got)
	}
}

func TestInOrder(t *testing.T) {
	tree := CreateTestTree()

	buffer := bytes.Buffer{}
	InOrder(tree, &buffer)

	got := buffer.String()
	want := "57429163810"

	if got != want {
		t.Errorf("in-order traversal failed, wanted %v got %v", want, got)
	}
}
