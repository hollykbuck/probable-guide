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

func isRed(x *Node) bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

func (t *Tree) get(key bst.Key) (interface{}, error) {
	if key == nil {
		return nil, bst.ErrInvalidArgument
	}
	return get(t.root, key)
}

func get(x *Node, key bst.Key) (interface{}, error) {
	for {
		if x != nil {
			break
		}
		cmp := key.CompareTo(x.key)
		if cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			return x.val, nil
		}
	}
	return nil, nil
}
