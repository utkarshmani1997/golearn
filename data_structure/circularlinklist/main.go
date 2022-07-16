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

type circularlinklist interface {
	add(string)
	addAtIndex(int, string)
	remove(int)
	get(int)
}

func (l *list) add(val string) {
	node := &node{
		val:  val,
		next: l.head,
	}
	if l.head == nil {
		l.head = node
		node.next = l.head
		l.length++
		return
	}
	next := l.head
	for next.next != l.head {
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
	l.add("b")
	l.add("c")
	l.add("d")
	l.add("e")
	l.add("f")
	l.remove(3)
	l.addAtIndex(5, "d")
	var next *node
	for next != l.head {
		if next == nil {
			next = l.head
		}
		fmt.Println(next.val, l.length)
		next = next.next
	}
}
