package treap

import (
	"cmp"
	"fmt"
	"math/rand"
)

type Treap[T cmp.Ordered] struct {
	root *Node[T]
}

type Node[T cmp.Ordered] struct {
	Priority int
	Key      T
	Left     *Node[T]
	Right    *Node[T]
}

func (t *Treap[T]) InOrderPrint() {
	inOrderPrint(t.root)
}

func inOrderPrint[T cmp.Ordered](n *Node[T]) {
	if n == nil {
		return
	}
	inOrderPrint(n.Left)
	fmt.Printf("p: %d, k: %v\n", n.Priority, n.Key)
	inOrderPrint(n.Right)
}

func (t *Treap[T]) Insert(key T) {
	t.root = insert(t.root, key)
}

func insert[T cmp.Ordered](root *Node[T], key T) *Node[T] {
	if root == nil {
		p := rand.Int()
		fmt.Printf("inserted p: %d, k: %v\n", p, key)
		return &Node[T]{Priority: p, Key: key}
	}

	if root.Key <= key {
		root.Right = insert(root.Right, key)
		if root.Priority < root.Right.Priority {
			root = rotateLeft(root)
		}
	} else {
		root.Left = insert(root.Left, key)
		if root.Priority < root.Left.Priority {
			root = rotateRight(root)
		}
	}

	return root
}

func rotateRight[T cmp.Ordered](n *Node[T]) *Node[T] {
	fmt.Printf("rotating right p: %d, k: %v\n", n.Priority, n.Key)
	x := n.Left
	n.Left = x.Right
	x.Right = n
	return x
}

func rotateLeft[T cmp.Ordered](n *Node[T]) *Node[T] {
	fmt.Printf("rotating left p: %d, k: %v\n", n.Priority, n.Key)
	x := n.Right
	n.Right = x.Left
	x.Left = n
	return x
}
