package main

import "fmt"

type node struct {
	val  string
	next *node
}

type list struct {
	length int
	head   *node
}

type linklist interface {
	add(string)
	addAtIndex(int, string)
	remove(int)
	get(int)
}

func (l *list) add(val string) {
	node := &node{
		val:  val,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		l.length++
		return
	}
	next := l.head
	for next.next != nil {
		next = next.next
	}
	next.next = node
	l.length++
}

func (l *list) addAtIndex(idx int, val string) {
	if idx == l.length {
		l.add(val)
		return
	}
	node := &node{
		val:  val,
		next: nil,
	}
	next := l.head
	prev := l.head
	for i := 0; i < idx; i++ {
		prev = next
		next = next.next
	}
	if next != nil {
		prev.next = node
		node.next = next
	}
}

func (l *list) remove(idx int) {
	next := l.head
	prev := l.head
	for i := 0; i < idx; i++ {
		prev = next
		next = next.next
	}
	if next != nil {
		prev.next = next.next
		l.length--
	}
}

func (l *list) get(idx int) *node {
	next := l.head
	for i := 0; i < idx; i++ {
		next = next.next
	}
	return next
}

func main() {
	l := &list{}
	l.add("a")
	head := add(nil, "a")
	add(head, "b")
	l.add("b")
	l.add("c")
	add(head, "c")
	l.add("d")
	add(head, "d")
	l.add("e")
	add(head, "e")
	l.add("f")
	add(head, "f")
	l.remove(3)
	l.addAtIndex(5, "d")
	next := l.head
	next1 := head
	i := 0
	for next != nil {
		fmt.Println(l.get(i).val, l.length)
		fmt.Println(next1.val)
		next = next.next
		next1 = next1.next
		i++
	}
}
func add(head *node, val string) *node {
	entry := &node{
		val:  val,
		next: nil,
	}
	if head == nil {
		head = entry
		return head
	}
	next := head
	prev := head
	for next != nil {
		prev = next
		next = next.next
	}
	prev.next = entry
	return head
}

func addAtIndex(idx int, val string) {
}

func remove(idx int) {
}

func get(idx int) {
}
