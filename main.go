package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_LEVEL = 16
)

type Node struct {
	forward []*Node
	key     interface{}
}

type SkipList struct {
	header *Node
	level  int
}

func newNode(level int, key interface{}) *Node {
	return &Node{
		forward: make([]*Node, level, level),
		key:     key,
	}
}

func newSkipList() *SkipList {
	return &SkipList{
		header: newNode(MAX_LEVEL, 0),
		level:  1,
	}
}

func (sl *SkipList) Insert(key interface{}) {
	update := make([]*Node, MAX_LEVEL)
	x := sl.header
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key.(int) < key.(int) {
			x = x.forward[i]
		}
		update[i] = x
	}
	rand.Seed(time.Now().UnixNano())
	level := rand.Intn(MAX_LEVEL-1) + 1
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.header
		}
		sl.level = level
	}
	x = newNode(level, key)
	for i := 0; i < level; i++ {
		x.forward[i] = update[i].forward[i]
		update[i].forward[i] = x
	}
}

func (sl *SkipList) Search(key interface{}) *Node {
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key.(int) < key.(int) {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		return x
	}
	return nil
}

func main() {
	sl := newSkipList()
	sl.Insert(1)
	sl.Insert(2)
	sl.Insert(3)
	sl.Insert(4)
	sl.Insert(5)
	sl.Insert(6)
	fmt.Println(sl.Search(4).key)
	fmt.Println(sl.Search(7))
}
