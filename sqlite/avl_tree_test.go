import (
	"testing"
)

type test struct {
	key   int
	value int
}

func TestAVLTree(t *testing.T) {
	tree := NewAVLTree()
	tree.Insert(test{1, 1})
	tree.Insert(test{2, 2})
	tree.Insert(test{3, 3})
	tree.Insert(test{4, 4})
	tree.Insert(test{5, 5})
	tree.Insert(test{6, 6})
	tree.Insert(test{7, 7})
	tree.Insert(test{8, 8})
	tree.Insert(test{9, 9})
	tree.Insert(test{10, 10})
	tree.Insert(test{11, 11})
	tree.Insert(test{12, 12})
	tree.Insert(test{13, 13})
	tree.Insert(test{14, 14})
	tree.Insert(test{15, 15})
	tree.Insert(test{16, 16})
	tree.Insert(test{17, 17})
	tree.Insert(test{18, 18})
	tree.Insert(test{19, 19})
	tree.Insert(test{20, 20})
	tree.Insert(test{21, 21})
	tree.Insert(test{22, 22})
	tree.Insert(test{23, 23})
	tree.Insert(test{24, 24})
	tree.Insert(test{25, 25})
	tree.Insert(test{26, 26})
	tree.Insert(test{27, 27})
	tree.Insert(test{28, 28})

	if tree.Size() != 28 {
		t.Error("tree size is not 28")
	}

	if tree.Len() != 28 {
		t.Error("tree len is not 28")
	}

	if tree.Empty() {
		t.Error("tree is empty")
	}

	if tree.Search(test{1, 1}) == nil {
		t.Error("tree search failed")
	}

	if tree.Search(test{29, 29}) != nil {
		t.Error("tree search failed")
	}

	if tree.Min() == nil {
		t.Error("tree min failed")
	}

	if tree.Max() == nil {
		t.Error("tree max failed")
	}

	if tree.Delete(test{1, 1}) == nil {
		t.Error("tree delete failed")
	}
}