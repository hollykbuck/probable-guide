package avl

import (
	"github.com/hollykbuck/probable-guide/bst"
)

// Node AVL 树的节点
type Node struct {
	key bst.Key
	val interface{}
	// height 树的高度。等于 1 + maxInt(height(left), height(right))
	height int
	// size 树的总大小。等于 1 + size(left) + size(right)
	size  int
	left  *Node
	right *Node
}

// Tree AVL 树数据结构
type Tree struct {
	root *Node
}

func (t *Tree) isEmpty() bool {
	return t.root == nil
}

func (t *Tree) size() int {
	return size(t.root)
}

func size(x *Node) int {
	if x == nil {
		return 0
	}
	return x.size
}

// height 树的高度
func (t *Tree) height() int {
	return height(t.root)
}

// height 子树的高度
func height(x *Node) int {
	if x == nil {
		return 0
	}
	return x.height
}

// get 在树中查询 key 对应的 value
func (t *Tree) get(key bst.Key) (interface{}, error) {
	if key == nil {
		return nil, bst.ErrInvalidArgument
	}
	x := get(t.root, key)
	if x == nil {
		return nil, nil
	}
	return x.val, nil
}

// get 从子树中查询 key 对应的 node
func get(x *Node, key bst.Key) *Node {
	if x == nil {
		return nil
	}
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		return get(x.left, key)
	} else if cmp > 0 {
		return get(x.right, key)
	} else {
		return x
	}
}

func (t *Tree) contains(key bst.Key) bool {
	return get(t.root, key) != nil
}

func (t *Tree) put(key bst.Key, val interface{}) (err error) {
	if key == nil {
		return bst.ErrInvalidArgument
	}
	if val == nil {
		err := t.delete(key)
		if err != nil {
			return err
		}
		return nil
	}
	t.root, err = put(t.root, key, val)
	return err
}

func (t *Tree) delete(key bst.Key) error {
	if key == nil {
		return bst.ErrInvalidArgument
	}
	if !t.contains(key) {
		return nil
	}
	t.root = deleteFromNode(t.root, key)
	return nil
}

// deleteFromNode 删除 key 对应的节点，并返回新树
func deleteFromNode(x *Node, key bst.Key) *Node {
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left = deleteFromNode(x.left, key)
	} else if cmp > 0 {
		x.right = deleteFromNode(x.right, key)
	} else {
		// 如果没有左孩子或者右孩子，可以直接删除，不需要平衡
		if x.left == nil {
			return x.right
		} else if x.right == nil {
			return x.left
		} else {
			// 如果有左孩子或者右孩子，首先删除，然后重平衡
			t := x
			x = min(x.right)
			x.right = deleteMin(t.right)
			x.left = t.left
		}
	}
	// 更新高度和大小
	x.size = 1 + size(x.left) + size(x.right)
	x.height = 1 + maxInt(height(x.left), height(x.right))
	// 删除一个节点只需要调整一次
	return balance(x)
}

// balance 重平衡算法。
// 重平衡不涉及递归。
func balance(x *Node) *Node {
	if balanceFactor(x) < -1 {
		// 右倾说明右子树很高
		// 先平衡右子树，再通过 rotateLeft 降低右子树高度
		if balanceFactor(x.right) > 0 {
			// 左倾说明左子树很高
			// 通过 rotateRight 降低左子树高度
			x.right = rotateRight(x.right)
		}
		x = rotateLeft(x)
	} else if balanceFactor(x) > 1 {
		// 左倾说明左子树很高
		// 先调整左子树，再通过 rotateRight 降低左子树高度
		if balanceFactor(x.left) < 0 {
			// 右倾说明右子树很高
			x.left = rotateLeft(x.left)
		}
		x = rotateRight(x)
	}
	return x
}

// rotateLeft 左旋转（逆时针旋转），并返回新的树。
// 左旋转只与 x 和 x.right 有关。效果是提升右子树。
// 将树根旋转为左子树。
func rotateLeft(x *Node) *Node {
	// 右子树提升为新的树根
	right := x.right
	// 接管 right 的子树
	x.right = right.left
	// 树根接管 x 的容量
	right.size = x.size
	x.size = 1 + size(x.left) + size(x.right)
	x.height = 1 + maxInt(height(x.left), height(x.right))
	right.height = 1 + maxInt(height(right.left), height(right.right))
	return right
}

// rotateRight 右旋转（顺时针旋转），并返回新的树。
// 右旋只与 x 和 x.left 有关。效果是提升左子树。
func rotateRight(x *Node) *Node {
	left := x.left
	x.left = left.right
	left.size = x.size
	x.size = 1 + size(x.left) + size(x.right)
	x.height = 1 + maxInt(height(x.left), height(x.right))
	left.height = 1 + maxInt(height(left.left), height(left.right))
	return left
}

// balanceFactor 平衡因子，左倾大于1, 右倾小于-1
func balanceFactor(x *Node) int {
	return height(x.left) - height(x.right)
}

func maxInt(i int, i2 int) int {
	if i > i2 {
		return i
	} else {
		return i2
	}
}

// deleteMin 删除最小元素
func (t *Tree) deleteMin(x *Node) error {
	if t.isEmpty() {
		return bst.ErrNoSuchElement
	}
	t.root = deleteMin(t.root)
	return nil
}

func deleteMin(x *Node) *Node {
	if x.left == nil {
		return x.right
	}
	x.left = deleteMin(x.left)
	x.size = 1 + size(x.left) + size(x.right)
	x.height = 1 + maxInt(height(x.left), height(x.right))
	return balance(x)
}

func min(x *Node) *Node {
	if x.left == nil {
		return x
	}
	return min(x.left)
}

func put(x *Node, key bst.Key, val interface{}) (_ *Node, err error) {
	if x == nil {
		return &Node{
			key:    key,
			val:    val,
			height: 0,
			size:   1,
			left:   nil,
			right:  nil,
		}, nil
	}
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left, err = put(x.left, key, val)
		if err != nil {
			return
		}
	} else if cmp > 0 {
		x.right, err = put(x.right, key, val)
		if err != nil {
			return
		}
	} else {
		x.val = val
		return x, nil
	}
	x.size = 1 + size(x.left) + size(x.right)
	x.height = 1 + maxInt(height(x.left), height(x.right))
	return balance(x), nil
}
