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
	n := &node{
		val:  val,
		next: nil,
	}
	if l.head == nil {
		l.head = n
		l.length++
		return
	}
	next := l.head
	for next.next != nil {
		next = next.next
	}
	next.next = n
	l.length++
}

func (l *list) addAtIndex(idx int, val string) {
	if idx == l.length {
		l.add(val)
		return
	}
	node := &node{
		val: val,
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
		l.length++
	}
}

func (l *list) remove(idx int) {
	next := l.head
	var prev *node
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
	fmt.Println(l.get(0), l.length)
	l.add("a")
	l.add("b")
	l.add("c")
	l.add("d")
	l.add("e")
	l.add("f")
	l.remove(3)
	l.addAtIndex(5, "d")
	next := l.head
	for next != nil {
		fmt.Println(next.val, l.length)
		next = next.next
	}
}
