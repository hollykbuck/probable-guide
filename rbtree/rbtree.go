package rbtree

import "github.com/hollykbuck/probable-guide/bst"

const (
	// RED 红表示 (2,4) B 树的关键字（额外）
	RED = true
	// BLACK 黑表示 (2,4) B 树的节点
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

func (t *Tree) delete(key bst.Key) error {
	if key == nil {
		return bst.ErrInvalidArgument
	}
	contains, err := t.contains(key)
	if err != nil {
		return err
	}
	// 如果不包含，直接结束
	if !contains {
		return nil
	}
	// 如果左右孩子都是黑，
	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = RED
	}
	t.root = deleteFromNode(t.root, key)
	if !t.isEmpty() {
		t.root.color = BLACK
	}
	return nil
}

func deleteFromNode(h *Node, key bst.Key) *Node {
	if key.CompareTo(h.key) < 0 {
		if !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}
	}
}

func moveRedLeft(h *Node) *Node {

}

func (t *Tree) contains(key bst.Key) (bool, error) {
	i, err := t.get(key)
	if err != nil {
		return false, err
	}
	return i != nil, nil
}

func (t *Tree) isEmpty() bool {
	return t.root == nil
}
