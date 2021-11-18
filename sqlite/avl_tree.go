package AVLTree

// AVLTree is a self-balancing binary search tree.
type AVLTree struct {
	root *node
	size int
}

type node struct {
	key   interface{}
	value interface{}
	left  *node
	right *node
	height int
}

type AVLTreeIterator struct {
	current *node
	tree    *AVLTree
}

type AVLTreeReverseIterator struct {
	current *node
	tree    *AVLTree
}

func (tree *AVLTree) Len() int {
	return tree.size
	
}

func (tree *AVLTree) Empty() bool {
	return tree.size == 0
}

func (tree *AVLTree) Clear() {
	tree.root = nil
	tree.size = 0
}

func (tree *AVLTree) Get(key interface{}) (interface{}, bool) {
	n := tree.root
	for n != nil {
		cmp := compare(key, n.key)
		switch {
		case cmp == 0:
			return n.value, true
		case cmp < 0:
			n = n.left
		default:
			n = n.right
		}
	}
	return nil, false
}

func (tree *AVLTree) Set(key interface{}, value interface{}) {
	n, p := tree.lookupOrCreate(key)
	n.value = value
	tree.rebalanceAfterSet(p)
}

func (tree *AVLTree) Remove(key interface{}) (value interface{}, ok bool) {
	n, p := tree.lookupParent(key)
	if n == nil {
		return nil, false
	}
	value = n.value
	tree.remove(n, p)
	return
}

func (tree *AVLTree) Has(key interface{}) bool {
	_, ok := tree.lookup(key)
	return ok
}

func (tree *AVLTree) First() (key, value interface{}) {
	n := tree.first()
	if n != nil {
		return n.key, n.value
	}
	return nil, nil
}

func (tree *AVLTree) Last() (key, value interface{}) {
	n := tree.last()
	if n != nil {
		return n.key, n.value
	}
	return nil, nil
}

func (tree *AVLTree) FirstKey() interface{} {
	n := tree.first()
	if n != nil {
		return n.key
	}
	return nil
}

func (tree *AVLTree) LastKey() interface{} {
	n := tree.last()
	if n != nil {
		return n.key
	}
	return nil
}

func (tree *AVLTree) Keys() []interface{} {
	keys := make([]interface{}, tree.size)
	i := 0
	tree.Ascend(func(key, value interface{}) bool {
		keys[i] = key
		i++
		return true
	})
	return keys
}

func (tree *AVLTree) Values() []interface{} {
	values := make([]interface{}, tree.size)
	i := 0
	tree.Ascend(func(key, value interface{}) bool {
		values[i] = value
		i++
		return true
	})
	return values
}

func (tree *AVLTree) Ascend(iterator ItemIterator) {
	var current *node
	for current = tree.first(); current != nil; current = tree.next(current) {
		if !iterator(current.key, current.value) {
			break
		}
	}
}

func (tree *AVLTree) Descend(iterator ItemIterator) {
	var current *node
	for current = tree.last(); current != nil; current = tree.prev(current) {
		if !iterator(current.key, current.value) {
			break
		}
	}
}

func (tree *AVLTree) AscendRange(from, to interface{}, iterator ItemIterator) {
	var current *node
	for current = tree.lowerBound(from); current != nil; current = tree.next(current) {
		if !iterator(current.key, current.value) {
			break
		}
		if compare(current.key, to) >= 0 {
			break
		}
	}
}

func (tree *AVLTree) DescendRange(from, to interface{}, iterator ItemIterator) {
	var current *node
	for current = tree.upperBound(from); current != nil; current = tree.prev(current) {
		if !iterator(current.key, current.value) {
			break
		}
		if compare(current.key, to) <= 0 {
			break
		}
	}
}

func (tree *AVLTree) Min() (key, value interface{}) {
	n := tree.min()
	if n != nil {
		return n.key, n.value
	}
	return nil, nil
}

func (tree *AVLTree) Max() (key, value interface{}) {
	n := tree.max()
	if n != nil {
		return n.key, n.value
	}
	return nil, nil
}

func (tree *AVLTree) Rebalance() {
	for current := tree.root; current != nil; current = current.left {
		tree.rebalanceAfterSet(current)
	}
}

func (tree *AVLTree) Iterator() *AVLTreeIterator {
	return &AVLTreeIterator{tree.first(), tree}
}

func (tree *AVLTree) ReverseIterator() *AVLTreeReverseIterator {
	return &AVLTreeReverseIterator{tree.last(), tree}
}

func (tree *AVLTree) doSet(n *node, key, value interface{}) {
	n.key = key
	n.value = value
}

func (tree *AVLTree) doRemove(n *node, parent *node) {
	if n.left != nil && n.right != nil {
		successor := n.right
		for successor.left != nil {
			successor = successor.left
		}
		tree.doSet(n, successor.key, successor.value)
		tree.doRemove(successor, n)
		return
	}

	if n.left != nil {
		tree.replace(n, n.left, parent)
	} else if n.right != nil {
		tree.replace(n, n.right, parent)
	} else {
		tree.replace(n, nil, parent)
	}
}

func (tree *AVLTree) doInsert(n *node, key, value interface{}) {
	n.key = key
	n.value = value
	tree.rebalanceAfterSet(n)
}

func (tree *AVLTree) doRotateLeft(n *node) {
	r := n.right
	tree.replace(n, r.left, r)
	n.right = r.left
	r.left = n
}

func (tree *AVLTree) doRotateRight(n *node) {
	l := n.left
	tree.replace(n, l.right, l)
	n.left = l.right
	l.right = n
}

func (tree *AVLTree) doDoubleRotateLeft(n *node) {
	tree.doRotateRight(n.right)
	tree.doRotateLeft(n)
}

func (tree *AVLTree) doDoubleRotateRight(n *node) {
	tree.doRotateLeft(n.left)
	tree.doRotateRight(n)
}

func (tree *AVLTree) replace(old, new *node, parent *node) {
	if parent == nil {
		tree.root = new
	} else if old == parent.left {
		parent.left = new
	} else {
		parent.right = new
	}
}

func (tree *AVLTree) rebalanceAfterSet(n *node) {
	for {
		if n == nil {
			return
		}

		if n.balance() == 2 {
			if n.left.balance() >= 0 {
				tree.doRotateLeft(n)
			} else {
				tree.doDoubleRotateLeft(n)
			}
		} else if n.balance() == -2 {
			if n.right.balance() <= 0 {
				tree.doRotateRight(n)
			} else {
				tree.doDoubleRotateRight(n)
			}
		}

		n = n.parent
	}
}
