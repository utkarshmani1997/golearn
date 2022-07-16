package main

import (
	"fmt"
)

type node struct {
	val  interface{}
	next *node
}

type list struct {
	length int
	rear   *node
	front  *node
}

type queue interface {
	enqueue(interface{})
	dequeue() interface{}
	peek() interface{}
	size() int
	display()
}

func New() queue {
	return new(list)
}

func (l *list) enqueue(val interface{}) {
	n := &node{
		val: val,
	}
	if l.rear == nil {
		l.rear = n
		l.front = n
		l.length++
		return
	}

	l.front.next = n
	l.front = n
	l.length++
}

func (l *list) dequeue() interface{} {
	if l.rear != nil {
		tmp := l.rear
		l.rear = l.rear.next
		l.length--
		return tmp.val
	}
	return nil
}

func (l *list) peek() interface{} {
	if l.rear != nil {
		return l.rear.val
	}
	return nil
}

func (l *list) size() int {
	return l.length
}

func (l *list) display() {
	for rear := l.rear; rear != nil; rear = l.rear {
		fmt.Println(l.dequeue(), l.peek(), l.size())
	}
}

func main() {
	l := New()
	l.display()
	l.enqueue(1)
	l.enqueue("s")
	l.enqueue("100")
	l.display()
}
