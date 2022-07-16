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

func search(root *Node, input int) bool {
	if root == nil {
		return false
	}
	if input > root.value {
		return search(root.right, input)
	}
	if input < root.value {
		return search(root.left, input)
	}

	return true
}

func main() {
	var root *Node
	scanner := bufio.NewScanner(os.Stdin)
menu:
	fmt.Println("Enter response: \n1. Insert\n2. Search\n3. Delete\n4. Inorder\n5. Preorder\n6. Postorder")
	for scanner.Scan() {
		stdIn := scanner.Text()
		switch stdIn {
		case "1":
			fmt.Println("Enter no of keys to be entered:")
			if scanner.Scan() {
				resp := scanner.Text()
				val, err := strconv.Atoi(resp)
				if err != nil {
					panic(err)
				}
				i := 0
				for i < val {
					scanner.Scan()
					resp := scanner.Text()
					val, err := strconv.Atoi(resp)
					if err != nil {
						panic(err)
					}
					root = buildTree(root, val)
					i++
				}
				goto menu
			}
		case "2":
			if root != nil {
				fmt.Println("\nEnter key to be searched: ")
				if scanner.Scan() {
					input := scanner.Text()
					val, err := strconv.Atoi(input)
					if err != nil {
						panic(err)
					}
					if search(root, val) {
						fmt.Println("Found")
					} else {
						fmt.Println("Not Found")
					}
				}
			} else {
				fmt.Println("Tree not found")
			}
			goto menu
		case "3":
			fmt.Println("Not implemented")
			goto menu
		case "4":
			fmt.Println("*********** INORDER TRAVERSAL *************")
			if root != nil {
				inorder(root)
			} else {
				fmt.Println("Tree not found")
			}
			goto menu
		case "5":
			fmt.Println("\n*********** PREORDER TRAVERSAL ************")
			if root != nil {
				preorder(root)
			} else {
				fmt.Println("Tree not found")
			}
			goto menu
		case "6":
			fmt.Println("\n*********** POSTORDER TRAVERSAL ***********")
			if root != nil {
				postorder(root)
			} else {
				fmt.Println("Tree not found")
			}
			goto menu
		default:
			panic("Invalid input")
		}
	}
}
