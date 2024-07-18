package treap

import (
	"cmp"
	"fmt"
	"math/rand"
	"strings"
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

func (t *Treap[T]) Repr(verbose bool) string {
	type pair struct {
		node  *Node[T]
		level int
	}

	lines := []string{}
	nodes := []pair{{node: t.root, level: 0}}

	for len(nodes) != 0 {
		var item pair
		item, nodes = nodes[len(nodes)-1], nodes[:len(nodes)-1]
		var s string
		if item.node != nil {
			if verbose {
				s = fmt.Sprintf("(p: %d, k: %v)", item.node.Priority, item.node.Key)
			} else {
				s = fmt.Sprintf("%v", item.node.Key)
			}
		} else {
			s = "Null"
		}
		indent := strings.Repeat(" ", item.level)
		lines = append(lines, fmt.Sprintf("%s%s", indent, s))

		if item.node != nil {
			nodes = append(
				nodes,
				pair{node: item.node.Right, level: item.level + 1},
				pair{node: item.node.Left, level: item.level + 1},
			)
		}
	}

	return strings.Join(lines, "\n")
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
	x := n.Left
	n.Left = x.Right
	x.Right = n
	return x
}

func rotateLeft[T cmp.Ordered](n *Node[T]) *Node[T] {
	x := n.Right
	n.Right = x.Left
	x.Left = n
	return x
}

func (t *Treap[T]) Delete(key T) {
	t.root = delete(t.root, key)
}

func delete[T cmp.Ordered](root *Node[T], targetKey T) *Node[T] {
	if root == nil {
		return nil
	}

	if root.Key < targetKey {
		root.Right = delete(root.Right, targetKey)
	} else if root.Key > targetKey {
		root.Left = delete(root.Left, targetKey)
	} else {
		if root.Right == nil {
			return root.Left
		}
		if root.Left == nil {
			return root.Right
		}

		if root.Left.Priority < root.Right.Priority {
			root = rotateLeft(root)
			root.Left = delete(root.Left, targetKey)
		} else {
			root = rotateRight(root)
			root.Right = delete(root.Right, targetKey)
		}
	}

	return root
}

func (t *Treap[T]) Find(key T) *Node[T] {
	return find(t.root, key)
}

func find[T cmp.Ordered](node *Node[T], key T) *Node[T] {
	if node == nil {
		return nil
	}

	if node.Key < key {
		return find(node.Right, key)
	}
	if node.Key > key {
		return find(node.Left, key)
	}

	return node
}
