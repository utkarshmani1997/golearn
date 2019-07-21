package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func buildTree(parent *Node, val int) *Node {
	if parent == nil {
		parent = &Node{val, nil, nil}
	} else if val < parent.value {
		parent.left = buildTree(parent.left, val)
	} else {
		parent.right = buildTree(parent.right, val)
	}

	return parent
}

// inorder traversal always prints in sorted order
// left->root->right
func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("%d, ", root.value)
	inorder(root.right)
}

// preorder traversal
// root->left->right
func preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d, ", root.value)
	preorder(root.left)
	preorder(root.right)
}

// postorder traversal
// left->right->root
func postorder(root *Node) {
	if root == nil {
		return
	}
	postorder(root.left)
	postorder(root.right)
	fmt.Printf("%d, ", root.value)
}

func main() {
	var root *Node
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter integers followed by enter button:")
	for scanner.Scan() {
		stdIn := scanner.Text()
		if stdIn == "" {
			break
		}
		val, err := strconv.Atoi(stdIn)
		if err != nil {
			panic(err)
		}
		root = buildTree(root, val)
	}

	fmt.Println("*********** INORDER TRAVERSAL *************")
	inorder(root)
	fmt.Println("\n*********** PREORDER TRAVERSAL ************")
	preorder(root)
	fmt.Println("\n*********** POSTORDER TRAVERSAL ***********")
	postorder(root)
	fmt.Println()
}
