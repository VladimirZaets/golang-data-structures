package unionfind

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(10)
	uf.Unify(0, 4)
	if uf.Size() == uf.Components() {
		t.Errorf("The component length shold be decreased after unify call, size %d, got, %d", uf.Size(), uf.Components())
	}
	uf.Unify(0, 3)
	if uf.Find(3) != 4 {
		t.Errorf("The component should have parent 4, got %d", uf.Find(3))
	}
	if uf.Find(5) != 5 {
		t.Errorf("The component should have parent 5, got %d", uf.Find(5))
	}
	if uf.Components() != 8 {
		t.Errorf("After 3 union operation should have 8 compinents, got %d", uf.Components())
	}
	if uf.ComponentSize(0) != 7 {
		t.Errorf("After 3 union operation, component size (sum of unified components) should be 7, got %d", uf.ComponentSize(0))
	}

}
