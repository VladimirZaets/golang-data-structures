package hashtable

import (
	"fmt"
	"testing"
)

func TestHashTableSeparateChaining(t *testing.T) {
	ht := NewHashTableSeparateChainint(1)
	for i := 0; i < 50; i++ {
		ht.Set(fmt.Sprintf("key-%d", i), fmt.Sprintf("val-%d", i))
	}
	for i := 0; i < 20; i += 2 {
		ht.Set(fmt.Sprintf("key-%d", i), fmt.Sprintf("val-%d-updated", i))
	}
	for i := 1; i < 50; i += 2 {
		if ht.Get(fmt.Sprintf("key-%d", i)) != fmt.Sprintf("val-%d", i) {
			t.Errorf("The value with key 'key-%d' should be equal 'val-%d'", i, i)
		}
	}
	for i := 0; i < 20; i += 2 {
		if ht.Get(fmt.Sprintf("key-%d", i)) != fmt.Sprintf("val-%d-updated", i) {
			t.Errorf("The value with key 'key-%d' should be equal 'val-%d-updated'", i, i)
		}
	}
	for i := 0; i < 30; i += 3 {
		ht.Remove(fmt.Sprintf("key-%d", i))
	}
	for i := 0; i < 30; i += 3 {
		if ht.Get(fmt.Sprintf("key-%d", i)) != nil {
			t.Errorf("The value with key 'key-%d' should be nil", i)
		}
	}

	fmt.Println("QQ")
}
