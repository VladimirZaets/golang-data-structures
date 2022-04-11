package fenwicktree

import (
	"reflect"
	"testing"
)

func TestFenwickTree(t *testing.T) {
	ft := NewFenwickTree([]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90})
	if !reflect.DeepEqual(ft.list, []int{0, 10, 30, 30, 100, 50, 110, 70, 360, 90}) {
		t.Error("The lists should be equal")
	}
	sum := ft.Sum(3, 6)
	if sum != 180 {
		t.Errorf("The sum should be %d, got %d", 180, sum)
	}
	ft.Add(3, 30)
	sum = ft.Sum(3, 6)
	if sum != 210 {
		t.Errorf("The sum should be %d, got %d", 210, sum)
	}
	ft.Set(4, 70)
	sum = ft.Sum(3, 6)
	if sum != 240 {
		t.Errorf("The sum should be %d, got %d", 240, sum)
	}
}
