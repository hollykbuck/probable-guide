package rbtree

import "github.com/hollykbuck/probable-guide/bst"

const (
	RED   = true
	BLACK = false
)

// Node 红黑树有额外的属性 color
type Node struct {
	key   bst.Key
	val   interface{}
	left  *Node
	right *Node
	color bool
	size  int
}

type Tree struct {
	root *Node
}
