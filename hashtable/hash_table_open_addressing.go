package hashtable

import (
	"fmt"
)

type ProbingInterface interface {
	Get(key interface{}, i int) int
}

type HashTableOpenAddressing struct {
	capacity      int
	maxLoadFactor float64
	threshold     int
	size          int
	list          []*HashTableItem
	probing       ProbingInterface
}

const tombstone = "tombstone"

func NewHashTableOpenAddressing(capacity int, probing ProbingInterface) *HashTableOpenAddressing {
	maxLoadFactor := 0.45
	return &HashTableOpenAddressing{
		capacity:      capacity,
		maxLoadFactor: maxLoadFactor,
		threshold:     int(maxLoadFactor * float64(capacity)),
		size:          0,
		list:          make([]*HashTableItem, capacity),
		probing:       probing,
	}
}

func (ht *HashTableOpenAddressing) hashKey(k interface{}) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

func (ht *HashTableOpenAddressing) getIndex(k interface{}) int {
	return ht.hashKey(k) % ht.capacity
}

func (ht *HashTableOpenAddressing) Set(k interface{}, v interface{}) {
	if ht.size >= ht.threshold {
		ht.resizeList()
	}
	index := ht.getIndex(k)
	i := 0

	for {
		if ht.list[index] == nil || ht.list[index].key == tombstone {
			ht.list[index] = &HashTableItem{key: k, value: v}
			ht.size++
			return
		}
		if ht.list[index].key == k {
			ht.list[index].value = v
			return
		}

		index = (index + ht.probing.Get(k, i)) % ht.capacity
		i++
	}
}

func (ht *HashTableOpenAddressing) Get(k interface{}) interface{} {
	index := ht.findIndex(k)
	if index != -1 {
		return ht.list[ht.findIndex(k)].value
	}
	return nil
}

func (ht *HashTableOpenAddressing) findIndex(k interface{}) int {
	index := ht.getIndex(k)
	i := 0
	for ht.list[index] != nil {
		if ht.list[index].key == k {
			return index
		}
		index = (index + ht.probing.Get(k, i)) % ht.capacity
		i++
	}
	return -1
}

func (ht *HashTableOpenAddressing) Remove(k interface{}) error {
	index := ht.findIndex(k)
	if index == -1 {
		return fmt.Errorf("element does not exist")
	}
	ht.list[index] = &HashTableItem{key: tombstone, value: nil}
	ht.size--
	return nil
}

func (ht *HashTableOpenAddressing) resizeList() {
	currentList := ht.list
	ht.capacity *= 2
	ht.threshold = int(ht.maxLoadFactor * float64(ht.capacity))
	ht.list = make([]*HashTableItem, ht.capacity)
	ht.size = 0
	for _, val := range currentList {
		if val == nil || val.key == tombstone {
			continue
		}
		ht.Set(val.key, val.value)
	}
}
