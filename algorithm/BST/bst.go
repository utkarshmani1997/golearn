package main

import "fmt"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

type Level struct {
	node  *BinaryTree
	Depth int
}

type stack struct {
	top   int
	level []Level
}

func New() *stack {
	return &stack{top: -1, level: make([]Level, 0)}
}

func (s *stack) Push(node *BinaryTree, depth int) {
	s.level = append(s.level, Level{node: node, Depth: depth})
	s.top++
}

func (s *stack) Pop() Level {
	if s.IsEmpty() {
		return Level{}
	}
	node := s.level[s.top]
	s.level = s.level[:len(s.level)-1]
	s.top--
	return node
}

func (s *stack) IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func NodeDepths(root *BinaryTree) int {
	sumOfDepth := 0
	st := New()
	st.Push(root, 0)
	for !st.IsEmpty() {
		level := st.Pop()
		node, depth := level.node, level.Depth
		if node == nil {
			continue
		}
		sumOfDepth += depth
		st.Push(node.Left, depth+1)
		st.Push(node.Right, depth+1)
	}
	return sumOfDepth
}

func BranchSums(root *BinaryTree) []int {
	sums := make([]int, 0)
	st := New()
	st.Push(root, 0)
	for !st.IsEmpty() {
		level := st.Pop()
		node, depth := level.node, level.Depth
		if node.Left == nil && node.Right == nil {
			sums = append(sums, depth+node.Value)
			continue
		}
		if node.Left != nil {
			st.Push(node.Left, depth+node.Value)
		}
		if node.Right != nil {
			st.Push(node.Right, depth+node.Value)
		}
	}

	return sums
}

func main() {
	/*`{
		"tree": {
		"nodes": [
	{"id": "1", "left": "2", "right": "3", "value": 1},
	{"id": "2", "left": "4", "right": "5", "value": 2},
	{"id": "3", "left": "6", "right": "7", "value": 3},
	{"id": "4", "left": "8", "right": "9", "value": 4},
	{"id": "5", "left": null, "right": null, "value": 5},
	{"id": "6", "left": null, "right": null, "value": 6},
	{"id": "7", "left": null, "right": null, "value": 7},
	{"id": "8", "left": null, "right": null, "value": 8},
	{"id": "9", "left": null, "right": null, "value": 9}
	],
	"root": "1"
	}
	}`
	*/
	tree := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value: 8,
					Left:  nil,
					Right: nil,
				},
				Right: &BinaryTree{
					Value: 9,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BinaryTree{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryTree{
				Value: 7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(NodeDepths(tree))
	fmt.Println(BranchSums(tree))
}
