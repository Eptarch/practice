package linkedlist

import (
	"errors"
)

var ErrEmptyList = errors.New("empty list")

type List struct {
    FirstNode *Node
    LastNode *Node
}

type Node struct {
    Val interface{}
    NextNode *Node
    PrevNode *Node
}

func NewList(args ...interface{}) *List {
    l := List{}
    for _, arg := range args {
        l.PushBack(arg)
    }
    return &l
}

func (n *Node) Next() *Node {
	return n.NextNode
}

func (n *Node) Prev() *Node {
    return n.PrevNode
}

func (l *List) PushFront(v interface{}) {
    var n = Node{Val: v}
    if l.FirstNode != nil {
        l.FirstNode.PrevNode = &n
        n.NextNode = l.FirstNode
    }
    l.FirstNode = &n
    if l.LastNode == nil {
        l.LastNode = &n
    }
}

func (l *List) PushBack(v interface{}) {
    var n = Node{Val: v}
    if l.LastNode != nil {
        l.LastNode.NextNode = &n
        n.PrevNode = l.LastNode
    }
    l.LastNode = &n
    if l.FirstNode == nil {
        l.FirstNode = &n
    }
}

func (l *List) PopFront() (interface{}, error) {
	if l.FirstNode == nil {
        return nil, ErrEmptyList
    }
    n := l.FirstNode
    l.FirstNode = l.FirstNode.NextNode
    if l.FirstNode == nil {
        l.LastNode = nil
    } else {
        l.FirstNode.PrevNode = nil
    }
    return n.Val, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.LastNode == nil {
        return nil, ErrEmptyList
    }
    n := l.LastNode
    l.LastNode = l.LastNode.PrevNode
    if l.LastNode == nil {
        l.FirstNode = nil
    } else {
        l.LastNode.NextNode = nil
    }
    return n.Val, nil
}

func (l *List) Reverse() {
    for e := l.LastNode; e != nil; e = e.NextNode {
		e.PrevNode, e.NextNode = e.NextNode, e.PrevNode
	}
	l.FirstNode, l.LastNode = l.LastNode, l.FirstNode
}

func (l *List) First() *Node {
	return l.FirstNode
}

func (l *List) Last() *Node {
	return l.LastNode
}
