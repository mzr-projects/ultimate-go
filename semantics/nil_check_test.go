package semantics

import (
	"testing"
)

func TestIntTree_Add(t *testing.T) {

}

func TestIntTree_Contains(t *testing.T) {
	tree := NewEmptyIntTree()
	tree = tree.Add(5)
	tree = tree.Add(3)
	tree = tree.Add(7)

	if !tree.Contains(7) {
		t.Error("Tree should contains added value.")
	}
}
