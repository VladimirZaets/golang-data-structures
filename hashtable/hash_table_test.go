package hashtable

import (
	"fmt"
	"testing"

	"github.com/VladimirZaets/godatastructures/hashtable/probing"
)

type hashTableTestType struct {
	htType   string
	instance HashTableInterface
}

func TestHashTable(t *testing.T) {
	sl := []*hashTableTestType{
		{htType: "separate-chaining", instance: NewHashTableSeparateChainint(1)},
		{htType: "linear-probing", instance: NewHashTableOpenAddressing(6, &probing.Linear{})},
		{htType: "quadratic-probing", instance: NewHashTableOpenAddressing(6, &probing.Quadratic{})},
		{htType: "double-hashing", instance: NewHashTableOpenAddressing(6, &probing.DoubleHashing{})},
	}
	for _, val := range sl {
		go testPerType(val.instance, val.htType, t)
	}
}

func testPerType(ht HashTableInterface, htType string, t *testing.T) {
	for i := 0; i < 50; i++ {
		ht.Set(fmt.Sprintf("key-%d", i), fmt.Sprintf("val-%d", i))
	}
	for i := 0; i < 20; i += 2 {
		ht.Set(fmt.Sprintf("key-%d", i), fmt.Sprintf("val-%d-updated", i))
	}
	for i := 1; i < 50; i += 2 {
		if ht.Get(fmt.Sprintf("key-%d", i)) != fmt.Sprintf("val-%d", i) {
			t.Errorf("Type %s: The value with key 'key-%d' should be equal 'val-%d'", htType, i, i)
		}
	}
	for i := 0; i < 20; i += 2 {
		if ht.Get(fmt.Sprintf("key-%d", i)) != fmt.Sprintf("val-%d-updated", i) {
			t.Errorf("Type %s: The value with key 'key-%d' should be equal 'val-%d-updated'", htType, i, i)
		}
	}
	for i := 0; i < 30; i += 3 {
		ht.Remove(fmt.Sprintf("key-%d", i))
	}
	for i := 0; i < 30; i += 3 {
		if ht.Get(fmt.Sprintf("key-%d", i)) != nil {
			t.Errorf("Type %s: The value with key 'key-%d' should be nil", htType, i)
		}
	}
}
