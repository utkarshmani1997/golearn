package main

import "fmt"

type node struct {
	val  interface{}
	prev *node
}

type list struct {
	length int
	top    *node
}

type stack interface {
	push(interface{})
	pop() interface{}
	peek() interface{}
	size() int
	show()
}

func New() stack {
	return new(list)
}

func (l *list) push(val interface{}) {
	l.top = &node{
		val:  val,
		prev: l.top,
	}
	l.length++
}

func (l *list) pop() interface{} {
	if l.top != nil {
		tmp := l.top
		l.top = l.top.prev
		l.length--
		return tmp.val
	}
	return nil
}

func (l *list) peek() interface{} {
	if l.top != nil {
		return l.top.val
	}
	return nil
}

func (l *list) size() int {
	return l.length
}

func (l *list) show() {
	for next := l.top; next != nil; next = l.top {
		val := l.pop()
		fmt.Println(val, l.length)
	}
}

func main() {
	l := New()
	l.push(1)
	l.push(2)
	l.push(3)
	l.show()
}
