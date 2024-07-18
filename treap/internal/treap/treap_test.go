// Tests success is non-deterministic because a Treap is a probabilistic binary
// search tree.

package treap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	tr := &Treap[string]{}
	assert.Equal(t, "Null", tr.Repr(false))

	tr.Insert("car")
	assert.Equal(t, "car\n Null\n Null", tr.Repr(false))

	tr.Insert("tree")
	assert.Equal(t, "car\n Null\n tree\n  Null\n  Null", tr.Repr(false))
}

func TestDelete(t *testing.T) {
	tr := &Treap[string]{}
	tr.Insert("car")
	tr.Insert("tree")
	tr.Insert("apple")

	tr.Delete("apple")
	assert.Equal(t, "car\n Null\n tree\n  Null\n  Null", tr.Repr(false))
}

func TestFind(t *testing.T) {
	tr := &Treap[string]{}
	tr.Insert("car")
	tr.Insert("tree")
	tr.Insert("apple")

	n := tr.Find("apple")
	assert.Equal(t, "apple", n.Key)
}
