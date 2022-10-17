package bst

import "fmt"

var (
	ErrInvalidArgument = &ErrorBST{s: "invalid argument"}
	ErrNoSuchElement   = &ErrorBST{s: "no such element"}
)

type ErrorBST struct {
	s string
}

func (b *ErrorBST) Error() string {
	return b.s
}

type Key interface {
	compareTo(Key) int
}

type Node struct {
	key   Key
	val   interface{}
	left  *Node
	right *Node
	size  int
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return &BST{root: nil}
}

func (b *BST) isEmpty() bool {
	return b.size() == 0
}

func (b *BST) size() int {
	return size(b.root)
}

func size(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.size
	}
}

func (b *BST) contains(key Key) (bool, error) {
	if key == nil {
		return false, ErrInvalidArgument
	}
	get, err := b.get(key)
	if err != nil {
		return false, err
	}
	return get != nil, nil
}

func (b *BST) get(key Key) (interface{}, error) {
	return getFromNode(b.root, key)
}

func getFromNode(x *Node, key Key) (interface{}, error) {
	if key == nil {
		return nil, ErrInvalidArgument
	}
	if x == nil {
		return nil, ErrInvalidArgument
	}
	to := key.compareTo(x.key)
	if to < 0 {
		return getFromNode(x.left, key)
	} else if to > 0 {
		return getFromNode(x.right, key)
	} else {
		return x.val, nil
	}
}

func (b *BST) put(key Key, val interface{}) error {
	if key == nil {
		return ErrInvalidArgument
	}
	if val == nil {
		return b.delete(key)
	}
	b.root = putToNode(b.root, key, val)
	return nil
}

// delete 移除指定的 key.
func (b *BST) delete(key Key) error {
	if key == nil {
		return fmt.Errorf("key == nil: %w", ErrInvalidArgument)
	}
	b.root = deleteFromNode(b.root, key)
	return nil
}

// deleteFromNode 删除节点 key 并返回新树
func deleteFromNode(x *Node, key Key) *Node {
	if x == nil {
		return nil
	}
	cmp := key.compareTo(x.key)
	// 左子树或者右子树和直接后继对换后删除原节点
	if cmp < 0 {
		// 在左子树上删除
		x.left = deleteFromNode(x.left, key)
	} else if cmp > 0 {
		// 在右子树上删除
		x.right = deleteFromNode(x.right, key)
	} else {
		// 如果右子树为空，该节点用左孩子替代，相当于直接移除了
		if x.right == nil {
			return x.left
		}
		// 如果左子树为空，该节点用右孩子替代，相当于直接移除了
		if x.left == nil {
			return x.right
		}
		// 如果左右子树都非空
		t := x
		// 找直接后继，用它替代
		x = min(t.right)
		// 移除直接后继
		x.right = deleteMinNode(t.right)
		// 左子树不变
		x.left = t.left
	}
	x.size = size(x.left) + size(x.right) + 1
	return x
}

func putToNode(x *Node, key Key, val interface{}) *Node {
	if x == nil {
		return &Node{
			key:  key,
			val:  val,
			size: 1,
		}
	}
	cmp := key.compareTo(x.key)
	if cmp < 0 {
		x.left = putToNode(x.left, key, val)
	} else if cmp > 0 {
		x.right = putToNode(x.right, key, val)
	} else {
		x.size = 1 + size(x.left) + size(x.right)
	}
	return x
}

func min(x *Node) *Node {
	if x.left == nil {
		return x
	} else {
		return min(x.left)
	}
}

func (b *BST) deleteMin() error {
	if b.isEmpty() {
		return fmt.Errorf("symbol table underflow: %w", ErrNoSuchElement)
	}
	b.root = deleteMinNode(b.root)
	return nil
}

// deleteMinNode 移除直接后继，返回新节点
func deleteMinNode(x *Node) *Node {
	// 如果已经到最左，将该节点用右子树代替，相当于删除
	if x.left == nil {
		return x.right
	}
	x.left = deleteMinNode(x.left)
	// 更新大小
	x.size = size(x.left) + size(x.right) + 1
	return x
}

func (b *BST) deleteMax() error {
	if b.isEmpty() {
		return ErrNoSuchElement
	}
	deleteMaxNode(b.root)
	return nil
}

func deleteMaxNode(x *Node) *Node {
	if x.right == nil {
		return x.left
	}
	x.right = deleteMaxNode(x.right)
	x.size = size(x.left) + size(x.right) + 1
	return x
}

// max 从根节点出发找最大值
func (b *BST) max() (Key, error) {
	if b.isEmpty() {
		return nil, ErrNoSuchElement
	}
	return max(b.root).key, nil
}

// max 递归在右子树中找最大值
func max(x *Node) *Node {
	if x.right == nil {
		return x
	} else {
		return max(x.right)
	}
}

func (b *BST) selectRank(rank int) (Key, error) {
	if rank < 0 || rank >= b.size() {
		return nil, ErrInvalidArgument
	}
	return selectRankFromNode(b.root, rank), nil
}

func selectRankFromNode(x *Node, rank int) Key {
	if x == nil {
		return nil
	}
	leftSize := size(x.left)
	if leftSize > rank {
		return selectRankFromNode(x.left, rank)
	} else if leftSize < rank {
		return selectRankFromNode(x.right, rank-leftSize-1)
	} else {
		return x.key
	}
}
