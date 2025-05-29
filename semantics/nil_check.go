package semantics

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Add(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}

	if val < it.val {
		it.left = it.left.Add(val)
	} else if val > it.val {
		it.right = it.right.Add(val)
	}

	return it
}

/*
Contains Here we don't modify the IntTree, but we need to check nil for it so we have to use pointer receiver to do this
*/
func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func NewEmptyIntTree() *IntTree {
	return nil
}
